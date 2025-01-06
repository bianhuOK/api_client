package iface

import (
	"time"

	"github.com/bianhuOK/api_client/internal/domain/model"
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
)

type DbRepository interface {
	ExecuteSql(sql string) ([]map[string]interface{}, error)
}

type DbFactory interface {
	GetDbRepository(model.DbConfig) (DbRepository, error)
}

type Cache interface {
	Get(key string) (*sql_template.SqlTemplate, bool)
	Set(key string, value *sql_template.SqlTemplate, ttl time.Duration)
}
