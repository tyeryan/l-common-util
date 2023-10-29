package database

import (
	"context"
	logutil "github.com/tyeryan/l-protocol/log"
	"gorm.io/gorm"
	"runtime/debug"
)

type Transactional interface {
	WithTransaction(ctx context.Context, txFun TransactionFunc) (err error)
}

type DBTransactional struct {
	gormConnector GORMConnector
}

type TransactionFunc func(tx *gorm.DB) error

func ProvideTransactional(gormConnector GORMConnector) Transactional {
	return &DBTransactional{gormConnector: gormConnector}
}

func (t *DBTransactional) WithTransaction(ctx context.Context, txFun TransactionFunc) (err error) {
	tx := t.gormConnector.GetDBWithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		log := logutil.GetLogger("Transactional")
		if p := recover(); p != nil {
			log.Errorw(ctx, "rollback transaction because of panic", "panic", p, "stacktrace", string(debug.Stack()))
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Errorw(ctx, "rollback transaction because of error", "error", err, "stacktrace", string(debug.Stack()))
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	return txFun(tx)
}
