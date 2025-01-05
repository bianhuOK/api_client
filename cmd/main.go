package main

import (
	"github.com/bianhuOK/api_client/app/app_sql_template"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/middleware/accesslog"
)

func main() {

	chassis.RegisterSchema("rest", &app_sql_template.ApiSqlController{})

	if err := chassis.Init(); err != nil {
		panic(err)
	}

	if err := chassis.Run(); err != nil {
		panic(err)
	}
}
