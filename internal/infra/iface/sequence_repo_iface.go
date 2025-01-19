package iface

import (
	"context"

	"github.com/bianhuOK/api_client/internal/domain/model"
)

// 序列仓储 - 基础设施，负责序列数据的持久化
type SequenceRepository interface {
	// 使用悲观锁获取序列（在事务中调用）
	GetSequenceForUpdate(ctx context.Context, name string) (*model.Sequence, error)
	// 更新序列值（在事务中调用）
	UpdateSequence(ctx context.Context, seq *model.Sequence) error
}
