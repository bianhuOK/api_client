package repo

import (
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/google/wire"
)

var SqlTemplateRepositorySet = wire.NewSet(
	NewRemoteApiTemplateRepository,
	wire.Bind(new(sql_template.TemplateRepository), new(*RemoteApiTemplateRepository)),
)
