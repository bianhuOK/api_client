//go:build wireinject
// +build wireinject

package app_sql_template

import (
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/internal/infra"
	"github.com/google/wire"
)

var SqlTemplateControllerSet = wire.NewSet(
	NewApiSqlController,
)

var SqlAppSet = wire.NewSet(
	infra.SqlTemplateInfraSet,
	sql_template.SqlTemplateServiceSet,
	SqlTemplateControllerSet,
)

func InitializeSqlApp() (*ApiSqlController, error) {
	wire.Build(SqlAppSet)
	return &ApiSqlController{}, nil
}
