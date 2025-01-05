package main

import (
	"github.com/bianhuOK/api_client/app/app_sql_template"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/middleware/accesslog"
)

func main() {

	app, err := app_sql_template.InitializeSqlApp()
	if err != nil {
		panic(err)
	}
	chassis.RegisterSchema("rest", app)

	if err := chassis.Init(); err != nil {
		panic(err)
	}

	if err := chassis.Run(); err != nil {
		panic(err)
	}
}
