package sql_template

import "time"

type TemplateRepository interface {
	GetTemplateById(id string) (*SqlTemplate, error)
}

type TemplateService interface {
	GetSqlTemplate(apiId string, params map[string]interface{}) (*SqlTemplate, error)
}

type Cache interface {
	Get(key string) (*SqlTemplate, bool)
	Set(key string, value *SqlTemplate, ttl time.Duration)
}

type RemoteAPI interface {
	FetchTemplate(id string) (*SqlTemplate, error)
}
