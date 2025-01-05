package app_sql_template

import (
	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

type ApiSqlController struct {
	SqlTemplateService domain.TemplateService
}

func NewApiSqlController(srv domain.TemplateService) *ApiSqlController {
	return &ApiSqlController{SqlTemplateService: srv}
}

func (c *ApiSqlController) QueryApiSql(b *rf.Context) {
	apiID := b.ReadPathParameter("api_id")
	var params map[string]interface{}
	err := b.ReadEntity(&params)
	if err != nil {
		b.WriteJSON(struct {
			Error string `json:"error"`
		}{Error: "Invalid request"}, "application/json")
		return
	}

	template, err := c.SqlTemplateService.GetSqlTemplate(apiID, params)
	if err != nil {
		b.WriteJSON(struct {
			Error string `json:"error"`
		}{Error: "Template not found"}, "application/json")
		return
	}

	b.WriteJSON(template, "application/json")
}

func (c *ApiSqlController) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: "POST", Path: "/api_sql_query/{api_id}", ResourceFunc: c.QueryApiSql,
			Returns: []*rf.Returns{{Code: 200}}},
	}
}
