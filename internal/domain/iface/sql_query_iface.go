package iface

import "github.com/bianhuOK/api_client/internal/domain/model"

type SqlQueryServiceIface interface {
	ExecuteSql(dc model.DbConfig, sql string) (*model.SqlQuery, error)
}
