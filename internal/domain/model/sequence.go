package model

import (
	"fmt"
	"sync"
)

type Sequence struct {
	Name             string  `gorm:"column:seq_name;type:varchar(64);not null;comment:'序列名称'"`
	CurrentValue     int64   `gorm:"column:current_value;type:bigint;not null;comment:'当前值'"`
	Step             int     `gorm:"column:step;type:int;not null;comment:'步长'"`
	MaxValue         int64   `gorm:"column:max_value;type:bigint;not null;comment:'最大值'"`
	PreloadThreshold float64 `gorm:"column:preload_threshold;type:float;not null;comment:'预加载阈值'"`
	Version          int     `gorm:"column:version;type:int;not null;default:1;comment:'版本号，用于乐观锁'"`
	CreatedAt        int     `gorm:"column:create_at;type:int;not null;comment:'创建时间(秒级时间戳)'"`
	ModifiedAt       int     `gorm:"column:modified_at;type:int;not null;comment:'修改时间(秒级时间戳)'"`
	Creator          string  `gorm:"column:creator;type:varchar(64);not null;default:'system';comment:'创建者'"`
	Modifier         string  `gorm:"column:modifier;type:varchar(64);not null;default:'system';comment:'修改者'"`
}

func (s *Sequence) TableName() string {
	return "sequence_table"
}

func NewSequence(name string, currentValue int64, step int, maxValue int64) *Sequence {
	return &Sequence{
		Name:         name,
		CurrentValue: currentValue,
		Step:         step,
		MaxValue:     maxValue,
	}
}

// SequenceRange 区间聚合根
type SequenceRange struct {
	Begin   int64
	End     int64
	Current int64
	mu      sync.Mutex
}

// NewSequenceRange 创建新的序列区间
func NewSequenceRange(begin, end int64) *SequenceRange {
	return &SequenceRange{
		Begin:   begin,
		End:     end,
		Current: begin,
	}
}

// Next 获取下一个值
func (r *SequenceRange) Next() (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.IsExhausted() {
		return 0, fmt.Errorf("sequence range exhausted")
	}

	value := r.Current
	r.Current++
	return value, nil
}

// IsExhausted 检查区间是否耗尽
func (r *SequenceRange) IsExhausted() bool {
	return r.Current >= r.End
}

type AsyncPreloadResult struct {
	Range *SequenceRange
	Err   error
}
