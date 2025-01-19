package repo

import (
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/internal/infra/configs"
	"github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	"github.com/google/wire"
)

var SqlTemplateRepositorySet = wire.NewSet(
	NewRemoteApiTemplateRepository,
	wire.Bind(new(sql_template.TemplateRepository), new(*RemoteApiTemplateRepository)),
)

var SequenceRepoSet = wire.NewSet(
	configs.LoadSequenceConfig,
	configs.LoadRangeConfig,
	configs.LoadDbOptionConfig,
	persistence.NewDBProvider,
	NewSequenceDbTransactionManager,
	NewSequenceRepository,
	wire.Bind(new(iface.SequenceRepository), new(*SequenceRepository)),
	wire.Bind(new(iface.Transaction), new(*SequenceDbTransactionManager)),
)

var DbFactoryImplSet = wire.NewSet(
	NewDbFactoryImpl,
	wire.Bind(new(iface.DbFactory), new(*DbFactoryImpl)),
)
