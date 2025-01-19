package persistence

import (
	"context"

	"gorm.io/gorm"
)

// 事务实现
type TransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

func (tm *TransactionManager) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, "tx", tx)
		return fn(ctx)
	})
}
