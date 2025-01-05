package app_sql_template

import (
	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
)

type ApiSqlController struct {
	SqlTemplateService domain.TemplateService
}

func NewApiSqlController(srv domain.TemplateService) *ApiSqlController {
	return &ApiSqlController{SqlTemplateService: srv}
}
