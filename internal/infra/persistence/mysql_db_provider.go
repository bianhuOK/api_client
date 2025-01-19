package persistence

import (
	"sync"
	"time"

	"github.com/bianhuOK/api_client/internal/infra/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// 全局连接池管理
	gormConnPool sync.Map
	// 连接池配置
	defaultPoolConfig = &configs.DatabaseOptionConfig{
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 5 * time.Hour,
		ConnMaxIdleTime: 1 * time.Hour,
	}
)

// DBProvider 数据库连接提供者
type DBProvider struct {
	config *configs.DatabaseOptionConfig
}

// NewDBProvider 创建数据库连接提供者
func NewDBProvider(config *configs.DatabaseOptionConfig) *DBProvider {
	if config == nil {
		config = defaultPoolConfig
	}
	return &DBProvider{config: config}
}

// GetConnection 获取数据库连接
func (p *DBProvider) GetConnection(dsn string) (*gorm.DB, error) {
	// 1. 检查连接池中是否存在
	if db, ok := gormConnPool.Load(dsn); ok {
		return db.(*gorm.DB), nil
	}

	// 2. 创建新连接
	db, err := p.createConnection(dsn)
	if err != nil {
		return nil, err
	}

	// 3. 存储到连接池
	gormConnPool.Store(dsn, db)
	return db, nil
}

// createConnection 创建新的数据库连接
func (p *DBProvider) createConnection(dsn string) (*gorm.DB, error) {
	// 1. GORM 配置
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	// 2. 创建连接
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	// 3. 获取底层 *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 4. 设置连接池参数
	sqlDB.SetMaxIdleConns(p.config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(p.config.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(p.config.ConnMaxIdleTime)

	return db, nil
}
