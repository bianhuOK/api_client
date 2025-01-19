package iface

import (
	"context"
)

type SequenceIface interface {
	TableName() string
	GetCurrentValue() int64
	SetCurrentValue(v int64) error
}

type SequenceRangeIface interface {
	Next() (int64, error)
	IsExhausted() bool
}

type SequenceServiceIface interface {
	// 获取下一个序列号
	NextValue(ctx context.Context) (int64, error)
	// 批量获取序列号
	// BatchNextValue(ctx context.Context, size int) ([]int64, error)
}
