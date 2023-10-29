package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/tyeryan/l-common-util/apm"
	"github.com/tyeryan/l-common-util/config"
	"github.com/tyeryan/l-common-util/secretclient"
	logutil "github.com/tyeryan/l-protocol/log"
	"go.elastic.co/apm/module/apmsql"
	"sync"
	"time"
)

var (
	WireSet = wire.NewSet(
		ProvideDatabaseConfig,
		ProvideSqlConnector,
		ProvideGORMConnector,
	)
	mutex = &sync.Mutex{}

	sqlConnectorInstance *sqlConnector
)

type DatabaseConfig struct {
	Type                     string `configstruct:"DatabaseConfig_Type" configdefault:""`
	Host                     string `configstruct:"DatabaseConfig_Host" configdefault:""`
	Port                     int    `configstruct:"DatabaseConfig_Port" configdefault:"0"`
	User                     string `configstruct:"DatabaseConfig_User" configdefault:""`
	Password                 string `configstruct:"DatabaseConfig_Password" configdefault:""`
	DatabaseName             string `configstruct:"DatabaseConfig_DatabaseName" configdefault:""`
	MaxIdleConns             int    `configstruct:"DatabaseConfig_MaxIdleConns" configdefault:"3"`
	MaxOpenConns             int    `configstruct:"DatabaseConfig_MaxOpenConns" configdefault:"5"`
	ServiceRole              string `configstruct:"DatabaseConfig_ServiceRole" configdefault:"youdb-service-role"`
	LogMode                  string `configstruct:"DatabaseConfig_LogMode" configdefault:"INFO"`
	EnableOnDemandCredential bool   `configstruct:"DatabaseConfig_EnableOnDemandCredential" configdefault:"false"`
	CredentialPath           string `configstruct:"DatabaseConfig_CredentialPath" configdefault:""`
}

func ProvideDatabaseConfig(ctx context.Context, configStore config.ConfigStore) (*DatabaseConfig, error) {
	log := logutil.GetLogger("database")

	dbConfig := &DatabaseConfig{}
	if err := configStore.GetConfig(dbConfig); err != nil {
		return nil, err
	}

	if len(dbConfig.Type) == 0 {
		return nil, errors.New("please provide db type")
	}
	if dbConfig.Type != "postgres" {
		return nil, errors.New(fmt.Sprintf("database type %s not suppported", dbConfig.Type))
	}
	if len(dbConfig.Host) == 0 {
		return nil, errors.New("please provide database host")
	}
	if dbConfig.Port == 0 {
		return nil, errors.New("please provide database port")
	}
	if len(dbConfig.DatabaseName) == 0 {
		return nil, errors.New("please provide database name")
	}
	if !dbConfig.EnableOnDemandCredential {
		if len(dbConfig.User) == 0 {
			return nil, errors.New("please provide database username for non on demand credential")
		}
		if len(dbConfig.Password) == 0 {
			return nil, errors.New("please provide database password for non on demand credential")
		}
	}

	log.Debugw(ctx, "database configured",
		"maxIdleConns", dbConfig.MaxIdleConns,
		"maxOpenConns", dbConfig.MaxOpenConns,
		"enableOnDemandCredential", dbConfig.EnableOnDemandCredential)

	return dbConfig, nil
}

type SqlConnector interface {
	GetDB() *sql.DB
}

type sqlConnector struct {
	db        *sql.DB
	cnf       *DatabaseConfig
	apmConfig *apm.ApmConfig

	refreshed chan bool
}

func (m *sqlConnector) GetDB() *sql.DB {
	return m.db
}

func (m *sqlConnector) connectToDatabase(ctx context.Context) error {
	log := logutil.GetLogger("database")
	// Connection args
	// see https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
	args := fmt.Sprintf(
		"sslmode=disable host=%s port=%d user=%s password='%s' dbname=%s",
		m.cnf.Host,
		m.cnf.Port,
		m.cnf.User,
		m.cnf.Password,
		m.cnf.DatabaseName,
	)

	db, err := m.openSqlConnection(args)
	if err != nil {
		log.Alerte(ctx, "connect to postgres Database failed", err, "args", args)
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		log.Alerte(ctx, "ping Database failed", err, "args", args)
		return err
	}

	log.Infow(ctx, "connect to postgres database successfully")

	// Max idle connections
	db.SetMaxIdleConns(m.cnf.MaxIdleConns)

	// Max open connections
	db.SetMaxOpenConns(m.cnf.MaxOpenConns)

	if m.db != nil {
		closeDB(ctx, m.db)
	}

	m.db = db
	return nil
}

func (m *sqlConnector) openSqlConnection(args string) (*sql.DB, error) {
	if m.apmConfig.Enable {
		return apmsql.Open(m.cnf.Type, args)
	} else {
		return sql.Open(m.cnf.Type, args)
	}
}

func (m *sqlConnector) refreshCredential(ctx context.Context, cnf *DatabaseConfig, secretStore secretclient.SecretClient) {
	log := logutil.GetLogger("database")
	dbCred := &DatabaseCredential{}
	duration, err := secretStore.Read(ctx, cnf.CredentialPath, dbCred)
	if err != nil {
		duration = time.Second * 5 // retry after 5 sec when error
		time.AfterFunc(duration, func() {
			m.refreshCredential(ctx, cnf, secretStore)
		})
		return
	}

	cnf.User = dbCred.UserName
	cnf.Password = dbCred.Password
	if err := m.connectToDatabase(ctx); err != nil {
		duration = time.Second * 5 // retry after 5 sec when error
		time.AfterFunc(duration, func() {
			m.refreshCredential(ctx, cnf, secretStore)
		})
		return
	}

	if duration > time.Minute*15 {
		duration -= time.Minute * 15 // refresh credential 15 mins before expiry
	}

	time.AfterFunc(duration, func() {
		m.refreshCredential(ctx, cnf, secretStore)
	})

	// non-blocking send
	select {
	case m.refreshed <- true:
	default:
	}

	log.Debugw(ctx, "refreshed db credential", "next", duration.String())
}

// ProvideSqlConnector SqlConnector provider
func ProvideSqlConnector(ctx context.Context, cnf *DatabaseConfig, secretStore secretclient.SecretClient, apmConfig *apm.ApmConfig) (SqlConnector, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if sqlConnectorInstance != nil {
		return sqlConnectorInstance, nil
	}

	sqlConnectorInstance = &sqlConnector{
		cnf:       cnf,
		apmConfig: apmConfig,
		refreshed: make(chan bool),
	}

	if cnf.EnableOnDemandCredential {
		sqlConnectorInstance.refreshCredential(ctx, cnf, secretStore)
	} else {
		if err := sqlConnectorInstance.connectToDatabase(ctx); err != nil {
			return nil, err
		}
	}

	return sqlConnectorInstance, nil
}

func closeDB(ctx context.Context, db *sql.DB) {
	log := logutil.GetLogger("database")

	inUse := db.Stats().InUse
	if inUse > 0 {
		log.Debugw(ctx, "still has conn in use", "inUse", inUse)
		time.AfterFunc(5*time.Second, func() {
			closeDB(ctx, db)
		})
	}

	log.Debugw(ctx, "going to close db", "dbStats", db.Stats())
	if err := db.Close(); err != nil {
		log.Errore(ctx, "fail to close db", err)
	}
}

type DatabaseCredential struct {
	UserName string `configstruct:"username"`
	Password string `configstruct:"password"`
}
