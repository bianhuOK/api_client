package service

import (
	"context"
	"sync"

	"github.com/bianhuOK/api_client/internal/domain/model"
	"github.com/bianhuOK/api_client/internal/infra/configs"
	"github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/bianhuOK/api_client/pkg/utils"
)

// rangeManager 管理序列区间
type RangeManager struct {
	repo        iface.SequenceRepository
	transaction iface.Transaction
	step        int64
}

func NewRangeManager(repo iface.SequenceRepository, tx iface.Transaction, rangeConfig *configs.RangeConfig) *RangeManager {
	return &RangeManager{
		repo:        repo,
		transaction: tx,
		step:        int64(rangeConfig.DefaultStep),
	}
}

func (rm *RangeManager) GetNextRange(ctx context.Context, name string) (*model.SequenceRange, error) {
	logger := utils.GetLogger()
	logger.Info("RangeManager GetNextRange", name)
	var sequenceRange *model.SequenceRange
	err := rm.transaction.ExecTx(ctx, func(ctx context.Context) error {
		// 1. 使用悲观锁获取序列
		seq, err := rm.repo.GetSequenceForUpdate(ctx, name)
		if err != nil {
			return err
		}

		// 2. 计算新区间
		newStart := seq.CurrentValue
		newEnd := newStart + rm.step
		seq.CurrentValue = newEnd

		// 4. 持久化更新
		if err := rm.repo.UpdateSequence(ctx, seq); err != nil {
			return err
		}

		// 5. 创建新区间
		sequenceRange = model.NewSequenceRange(newStart, newEnd)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return sequenceRange, nil
}

func (rm *RangeManager) IsExhausted(r *model.SequenceRange) bool {

	if r.Current >= r.End {
		return true
	} else {
		return false
	}
}

type PreloadManager struct {
	rangeManager     *RangeManager
	name             string
	preLoadThreshold float64
	mu               sync.Mutex
}

func NewPreloadManager(rm *RangeManager, rangeConfig *configs.RangeConfig) *PreloadManager {
	return &PreloadManager{
		rangeManager:     rm,
		name:             rangeConfig.Name,
		preLoadThreshold: rangeConfig.PreloadThreshold,
	}
}

func (pm *PreloadManager) StartPreload(ctx context.Context) (*model.AsyncPreloadResult, error) {
	resultChan := make(chan *model.AsyncPreloadResult, 1)

	go func() {
		range_, err := pm.rangeManager.GetNextRange(ctx, pm.name)
		resultChan <- &model.AsyncPreloadResult{
			Range: range_,
			Err:   err,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-resultChan:
		return result, nil
	}
}

func (pm *PreloadManager) ShouldPreload(currentRange *model.SequenceRange) bool {
	if currentRange == nil {
		return false
	}

	totalRange := float64(currentRange.End - currentRange.Begin)
	used := float64(currentRange.Current - currentRange.Begin)
	usedRatio := used / totalRange

	return usedRatio >= pm.preLoadThreshold
}

func (pm *PreloadManager) HandlePreloadResult(result *model.AsyncPreloadResult) {
	// 处理预加载结果的逻辑
	// 可以添加监控指标或日志
}
