package persistence

import (
	"github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/google/wire"
)

var SqlTemplateLocalCacheSet = wire.NewSet(
	ProviderSqlLocalCacheConfig,
	NewSqlLocalCache,
	wire.Bind(new(iface.Cache), new(*SqlLocalCache)),
)

var SqlQueryMysqlSet = wire.NewSet()
