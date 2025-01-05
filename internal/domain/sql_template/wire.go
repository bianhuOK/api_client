package sql_template

import (
	"github.com/google/wire"
)

var SqlTemplateServiceSet = wire.NewSet(
	NewSqlTemplateService,
	wire.Bind(new(TemplateService), new(*SqlTemplateService)),
)
