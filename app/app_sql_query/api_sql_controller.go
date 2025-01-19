package app_sql_query

import (
	"runtime/debug"

	"github.com/bianhuOK/api_client/internal/domain/iface"
	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/pkg/utils"
	"github.com/go-chassis/go-chassis/v2/pkg/metrics"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

type ApiSqlController struct {
	SqlTemplateService domain.TemplateService
	SqlQueryService    iface.SqlQueryServiceIface
}

func NewApiSqlController(sqlTemplateSrv domain.TemplateService, sqlQuerySrv iface.SqlQueryServiceIface) *ApiSqlController {

	return &ApiSqlController{
		SqlTemplateService: sqlTemplateSrv,
		SqlQueryService:    sqlQuerySrv,
	}
}

func (c *ApiSqlController) QueryApiSql(b *rf.Context) {
	logger := utils.GetLogger()
	// 记录请求
	metrics.CounterAdd("request_counter", 1, map[string]string{
		"method":   b.ReadRequest().Method,
		"endpoint": b.ReadRequest().URL.Path,
	})

	// 延迟执行的函数，用于捕获并处理panic
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
	// 根据api_id和params获取SQL模板
	template, err := c.SqlTemplateService.GetSqlTemplate(apiID, params)
	if err != nil {
		b.WriteJSON(struct {
			Error string `json:"error"`
		}{Error: "Template not found"}, "application/json")
		return
	}

	// 使用SQL模板和参数执行查询
	results, err := c.SqlQueryService.ExecuteSql(template.DbConfig, string(template.TemplateContent))
	if err != nil {
		b.WriteJSON(struct {
			Error string `json:"error"`
		}{Error: "Failed to execute query"}, "application/json")
		return
	}

	b.WriteJSON(results, "application/json")
}

func (c *ApiSqlController) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: "POST", Path: "/api_sql_query/{api_id}", ResourceFunc: c.QueryApiSql,
			Returns: []*rf.Returns{{Code: 200}}},
	}
}
