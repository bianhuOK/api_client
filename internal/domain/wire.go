package domain

import (
	"github.com/bianhuOK/api_client/internal/domain/iface"
	"github.com/bianhuOK/api_client/internal/domain/service"
	"github.com/google/wire"
)

var SqlQuerySet = wire.NewSet(
	service.NewSqlQueryService,
	wire.Bind(new(iface.SqlQueryServiceIface), new(*service.SqlQueryService)),
)
