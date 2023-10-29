package database

import (
	"context"
	"fmt"
	"github.com/tyeryan/l-common-util/apm"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	gormMutex    = &sync.Mutex{}
	gormInstance *gormConnector
)

func ProvideGORMConnector(ctx context.Context, cnf *DatabaseConfig, apmCnf *apm.ApmConfig, dbConnector SqlConnector) (GORMConnector, error) {
	gormMutex.Lock()
	defer gormMutex.Unlock()

	if gormInstance != nil {
		return gormInstance, nil
	}

	gormInstance = &gormConnector{
		sqlConnector: dbConnector.(*sqlConnector),
		cnf:          cnf,
		apmCnf:       apmCnf,
	}
	gormInstance.connect()

	go gormInstance.enableDynamicCredential()
	return gormInstance, nil
}

type GORMConnector interface {
	GetDB() *gorm.DB
	GetDBWithContext(ctx context.Context) *gorm.DB
}

type gormConnector struct {
	db           *gorm.DB
	cnf          *DatabaseConfig
	apmCnf       *apm.ApmConfig
	sqlConnector *sqlConnector
}

func (c *gormConnector) GetDB() *gorm.DB {
	return c.db
}

// propagate context to db connection
// apm requires this to be able to track db query time
func (c *gormConnector) GetDBWithContext(ctx context.Context) *gorm.DB {
	if c.apmCnf.Enable {
		return c.db.WithContext(ctx)
	} else {
		return c.GetDB()
	}
}

func (c *gormConnector) connect() {
	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Singapore",
		c.cnf.Host,
		c.cnf.User,
		c.cnf.Password,
		c.cnf.DatabaseName,
		c.cnf.Port,
	)

	db, err := c.openGormConnection(args)
	if err != nil {
		panic("fail to connect, gorm sucks!")
	}

	db.Logger.LogMode(LogLevel[c.cnf.LogMode])
	c.db = db
}

func (c *gormConnector) openGormConnection(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (c *gormConnector) enableDynamicCredential() {
	for {
		<-c.sqlConnector.refreshed
		c.connect()
	}
}

var LogLevel = map[string]logger.LogLevel{
	"INFO":   logger.Info,
	"WARN":   logger.Warn,
	"ERROR":  logger.Error,
	"SILENT": logger.Silent,
}
