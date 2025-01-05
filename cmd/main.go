package main

import (
	"github.com/bianhuOK/api_client/internal/demo"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/middleware/accesslog"
)

func main() {

	chassis.RegisterSchema("rest", &demo.RestFulHello{})

	if err := chassis.Init(); err != nil {
		panic(err)
	}

	if err := chassis.Run(); err != nil {
		panic(err)
	}
}
