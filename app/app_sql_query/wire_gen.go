// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app_sql_query

import (
	"github.com/bianhuOK/api_client/internal/domain"
	"github.com/bianhuOK/api_client/internal/domain/service"
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/internal/infra"
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	"github.com/bianhuOK/api_client/internal/infra/remote_api"
	"github.com/bianhuOK/api_client/internal/infra/repo"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeSqlApp() (*ApiSqlController, error) {
	localCacheConfig := persistence.ProviderSqlLocalCacheConfig()
	sqlLocalCache := persistence.NewSqlLocalCache(localCacheConfig)
	mockRemoteAPI := remoteapi.NewMockRemoteAPI()
	remoteApiTemplateRepository := repo.NewRemoteApiTemplateRepository(sqlLocalCache, mockRemoteAPI)
	sqlTemplateService := sql_template.NewSqlTemplateService(remoteApiTemplateRepository)
	dbFactoryImpl := repo.NewDbFactoryImpl()
	sqlQueryService := service.NewSqlQueryService(dbFactoryImpl)
	apiSqlController := NewApiSqlController(sqlTemplateService, sqlQueryService)
	return apiSqlController, nil
}

func InitializeGrpcSqlApp() (*GrpcSqlController, error) {
	localCacheConfig := persistence.ProviderSqlLocalCacheConfig()
	sqlLocalCache := persistence.NewSqlLocalCache(localCacheConfig)
	mockRemoteAPI := remoteapi.NewMockRemoteAPI()
	remoteApiTemplateRepository := repo.NewRemoteApiTemplateRepository(sqlLocalCache, mockRemoteAPI)
	sqlTemplateService := sql_template.NewSqlTemplateService(remoteApiTemplateRepository)
	dbFactoryImpl := repo.NewDbFactoryImpl()
	sqlQueryService := service.NewSqlQueryService(dbFactoryImpl)
	grpcSqlController := NewGrpcSqlController(sqlTemplateService, sqlQueryService)
	return grpcSqlController, nil
}

// wire.go:

var SqlTemplateControllerSet = wire.NewSet(
	NewApiSqlController,
)

var SqlTemplateGrpcControllerSet = wire.NewSet(
	NewGrpcSqlController,
)

var SqlAppCommonSet = wire.NewSet(infra.MockSqlTemplateInfraSet, infra.SqlQueryInfraSet, sql_template.SqlTemplateServiceSet, domain.SqlQuerySet)

var SqlAppSet = wire.NewSet(
	SqlAppCommonSet,
	SqlTemplateControllerSet,
)

var SqlGrpcAppSet = wire.NewSet(
	SqlAppCommonSet,
	SqlTemplateGrpcControllerSet,
)
