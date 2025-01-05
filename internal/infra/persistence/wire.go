package persistence

import (
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/google/wire"
)

var SqlTemplateLocalCacheSet = wire.NewSet(
	ProviderSqlLocalCacheConfig,
	NewSqlLocalCache,
	wire.Bind(new(sql_template.Cache), new(*SqlLocalCache)),
)
