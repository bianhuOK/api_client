// config/config.go
package configs

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config 序列服务配置
type SequenceConfig struct {
	DatabaseConfig       DatabaseConfig       `yaml:"database"`
	DatabaseOptionConfig DatabaseOptionConfig `yaml:"databaseConfig"`
	Range                RangeConfig          `yaml:"range"`
}

// RangeConfig 序列区间配置
type RangeConfig struct {
	Name             string        `yaml:"name"`
	DefaultStep      int           `yaml:"defaultStep"`
	MaxValue         int64         `yaml:"maxValue"`
	PreloadThreshold float64       `yaml:"preloadThreshold"`
	MaxRetries       int           `yaml:"maxRetries"`
	RetryDelay       time.Duration `yaml:"retryDelay"`
}

// LoadConfig 加载配置
func LoadSequenceConfig() (*SequenceConfig, error) {
	// 1. 确定配置文件路径
	configPath := getConfigPath()

	// 2. 读取配置文件
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 3. 解析配置
	config := &SequenceConfig{}
	if err := yaml.Unmarshal(configFile, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// 4. 验证配置
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return config, nil
}

func LoadDbOptionConfig() (*DatabaseOptionConfig, error) {
	config, err := LoadSequenceConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read seq config: %w", err)
	}
	return &config.DatabaseOptionConfig, nil
}

func LoadRangeConfig() (*RangeConfig, error) {
	config, err := LoadSequenceConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read seq config: %w", err)
	}
	return &config.Range, nil
}

// getConfigPath 获取配置文件路径
func getConfigPath() string {
	// 优先使用环境变量
	os.Setenv("SEQUENCE_CONFIG_PATH", "/Users/bianniu/GolandProjects/api_client/internal/infra/configs/sequence.local.yaml")
	if path := os.Getenv("SEQUENCE_CONFIG_PATH"); path != "" {
		return path
	}

	// 默认配置文件路径
	env := os.Getenv("SEQUENCE_ENV")
	if env == "" {
		env = "local"
	}

	return fmt.Sprintf("sequence.%s.yaml", env)
}

// validate 验证配置
func (c *SequenceConfig) validate() error {
	// 验证数据库配置
	db := c.DatabaseConfig
	if db.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if db.Port == 0 {
		return fmt.Errorf("database port is required")
	}
	if db.Username == "" {
		return fmt.Errorf("database username is required")
	}
	if db.Database == "" {
		return fmt.Errorf("database name is required")
	}

	// 验证数据库连接池配置
	dbConfig := c.DatabaseOptionConfig
	if dbConfig.MaxIdleConns <= 0 {
		return fmt.Errorf("maxIdleConns must be positive")
	}
	if dbConfig.MaxOpenConns <= 0 {
		return fmt.Errorf("maxOpenConns must be positive")
	}
	if dbConfig.MaxOpenConns < dbConfig.MaxIdleConns {
		return fmt.Errorf("maxOpenConns must be greater than or equal to maxIdleConns")
	}

	// 验证序列配置
	r := c.Range
	if r.DefaultStep <= 0 {
		return fmt.Errorf("default step must be positive")
	}
	if r.MaxValue <= 0 {
		return fmt.Errorf("max value must be positive")
	}
	if r.PreloadThreshold <= 0 || r.PreloadThreshold >= 1 {
		return fmt.Errorf("preload threshold must be between 0 and 1")
	}
	if r.MaxRetries <= 0 {
		return fmt.Errorf("max retries must be positive")
	}

	return nil
}
