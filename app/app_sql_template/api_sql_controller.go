package app_sql_template

import (
	"runtime/debug"

	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/pkg/utils"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

type ApiSqlController struct {
	SqlTemplateService domain.TemplateService
}

func NewApiSqlController(srv domain.TemplateService) *ApiSqlController {
	return &ApiSqlController{SqlTemplateService: srv}
}

func (c *ApiSqlController) QueryApiSql(b *rf.Context) {
	logger := utils.GetLogger()

	defer func() {
		if err := recover(); err != nil {
			logger.WithFields(map[string]interface{}{
				"panic": err,
				"stack": string(debug.Stack()),
			}).Error("handle request panic")
			b.WriteJSON(struct {
				Error string `json:"error"`
			}{Error: "Internal server error"}, "application/json")
		}
	}()

	apiID := b.ReadPathParameter("api_id")
	var params map[string]interface{}
	err := b.ReadEntity(&params)
	if err != nil {
		b.WriteJSON(struct {
			Error string `json:"error"`
		}{Error: "Invalid request"}, "application/json")
		return
	}
	logger.Info("QueryApiSql", "api_id: ", apiID, ", params: ", params)
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
