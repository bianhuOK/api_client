package sql_template

type TemplateRepository interface {
	GetTemplateById(id string) (*SqlTemplate, error)
}

type TemplateService interface {
	GetSqlTemplate(apiId string, params map[string]interface{}) (*SqlTemplate, error)
}
