package repo

import (
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/google/wire"
)

var SqlTemplateRepositorySet = wire.NewSet(
	NewRemoteApiTemplateRepository,
	wire.Bind(new(sql_template.TemplateRepository), new(*RemoteApiTemplateRepository)),
)

var DbFactoryImplSet = wire.NewSet(
	NewDbFactoryImpl,
	wire.Bind(new(iface.DbFactory), new(*DbFactoryImpl)),
)
