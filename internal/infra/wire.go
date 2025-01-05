package infra

import (
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	remoteapi "github.com/bianhuOK/api_client/internal/infra/remote_api"
	"github.com/bianhuOK/api_client/internal/infra/repo"
	"github.com/google/wire"
)

var SqlTemplateInfraSet = wire.NewSet(
	persistence.SqlTemplateLocalCacheSet,
	remoteapi.RemoteApiSet,
	repo.SqlTemplateRepositorySet,
)

var MockSqlTemplateInfraSet = wire.NewSet(
	persistence.SqlTemplateLocalCacheSet,
	remoteapi.MockRemoteApiSet,
	repo.SqlTemplateRepositorySet,
)
