package iface

import "github.com/bianhuOK/api_client/internal/domain/sql_template"

type RemoteAPI interface {
	FetchTemplate(id string) (*sql_template.SqlTemplate, error)
}
