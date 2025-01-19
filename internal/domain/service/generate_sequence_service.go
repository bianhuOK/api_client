package service

import (
	"context"
	"sync"

	"github.com/bianhuOK/api_client/internal/domain/model"
	"github.com/bianhuOK/api_client/internal/infra/configs"
)

type SeqGenerator struct {
	name         string
	rangeManager *RangeManager
	preloadMgr   *PreloadManager
	currentRange *model.SequenceRange
	nextRange    *model.SequenceRange
	mu           sync.Mutex
}

func NewSeqGenerator(rangeConfig *configs.RangeConfig, rangeManager *RangeManager, preloadMgr *PreloadManager) *SeqGenerator {
	return &SeqGenerator{
		name:         rangeConfig.Name,
		rangeManager: rangeManager,
		preloadMgr:   preloadMgr,
	}
}

func (s *SeqGenerator) NextValue(ctx context.Context) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 1. 检查当前区间
	if s.currentRange != nil && !s.rangeManager.IsExhausted(s.currentRange) {
		// 检查是否需要预加载
		if s.nextRange == nil && s.preloadMgr.ShouldPreload(s.currentRange) {
			go s.triggerPreload(context.Background())
		}
		return s.currentRange.Next()
	}

	// 2. 使用预加载的区间
	if s.nextRange != nil {
		s.currentRange = s.nextRange
		s.nextRange = nil
		return s.currentRange.Next()
	}

	// 3. 获取新区间
	newRange, err := s.rangeManager.GetNextRange(ctx, s.name)
	if err != nil {
		return 0, err
	}

	s.currentRange = newRange
	return s.currentRange.Next()
}

func (s *SeqGenerator) triggerPreload(ctx context.Context) {
	result, err := s.preloadMgr.StartPreload(ctx)
	if err != nil {
		// 处理错误，可以记录日志
		return
	}

	s.mu.Lock()
	defer s.mu.Lock()

	if s.nextRange == nil {
		s.nextRange = result.Range
	}

	s.preloadMgr.HandlePreloadResult(result)
}
