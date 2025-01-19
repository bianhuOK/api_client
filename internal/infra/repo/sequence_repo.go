package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/bianhuOK/api_client/internal/domain/model"
	"github.com/bianhuOK/api_client/internal/infra/configs"
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SequenceRepository struct {
	db *gorm.DB
}

func NewSequenceRepository(provider *persistence.DBProvider, config *configs.SequenceConfig) (*SequenceRepository, error) {
	dsn := config.DatabaseConfig.GetDSN()
	db, err := provider.GetConnection(dsn)
	if err != nil {
		return nil, err
	}

	return &SequenceRepository{db: db}, err
}

func (r *SequenceRepository) GetSequenceForUpdate(ctx context.Context, name string) (*model.Sequence, error) {
	tx := r.getTransaction(ctx)
	var seq model.Sequence

	// 使用 SELECT FOR UPDATE 加悲观锁
	err := tx.Table(seq.TableName()).Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("seq_name = ?", name).
		First(&seq).Error
	if err != nil {
		return nil, err
	}

	return &seq, nil
}

func (r *SequenceRepository) UpdateSequence(ctx context.Context, seq *model.Sequence) error {
	tx := r.getTransaction(ctx)

	// 使用乐观锁更新
	// todo modifier
	result := tx.Table(seq.TableName()).
		Where("seq_name = ? AND version = ?", seq.Name, seq.Version).
		Updates(map[string]interface{}{
			"current_value": seq.CurrentValue,
			"version":       seq.Version + 1,
			"modified_at":   time.Now().Unix(),
		})

	if result.RowsAffected == 0 {
		return fmt.Errorf("sequence was updated by another process")
	}

	return result.Error
}

// 从上下文获取事务对象
func (r *SequenceRepository) getTransaction(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		return r.db
	}
	return tx
}

type SequenceDbTransactionManager struct {
	*persistence.TransactionManager
}

func NewSequenceDbTransactionManager(provider *persistence.DBProvider, config *configs.SequenceConfig) (*SequenceDbTransactionManager, error) {
	dsn := config.DatabaseConfig.GetDSN()
	db, err := provider.GetConnection(dsn)
	if err != nil {
		return nil, err
	}
	tm := persistence.NewTransactionManager(db)
	return &SequenceDbTransactionManager{
		tm,
	}, nil

}
