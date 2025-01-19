package main

import (
	"os"

	"github.com/bianhuOK/api_client/app/app_sequence_get/schemas"
	"github.com/bianhuOK/api_client/app/app_sql_query"
	pb "github.com/bianhuOK/api_client/app/app_sql_query/proto" // 导入生成的proto
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	_ "github.com/go-chassis/go-chassis/v2/middleware/accesslog"
	_ "github.com/go-chassis/go-chassis/v2/middleware/monitoring"
	"github.com/go-chassis/go-chassis/v2/pkg/metrics"
	_ "github.com/go-chassis/go-chassis/v2/pkg/metrics"
)

func main() {
	os.Setenv("CHASSIS_HOME", "/Users/bianniu/GolandProjects/api_client")
	// app, err := app_sql_query.InitializeSqlApp()
	appSequence, err := schemas.InitializeSequenceApp()
	appGrpc, err := app_sql_query.InitializeGrpcSqlApp()
	if err != nil {
		panic(err)
	}

	// chassis.RegisterSchema("rest", app)
	chassis.RegisterSchema("rest", appSequence)
	chassis.RegisterSchema("grpc", appGrpc, server.WithRPCServiceDesc(&pb.SqlQueryService_ServiceDesc))

	if err := chassis.Init(); err != nil {
		panic(err)
	}

	//
	err = metrics.Init()
	if err != nil {
		panic(err)
	}

	err = metrics.CreateCounter(metrics.CounterOpts{
		Name:   "request_counter",
		Help:   "A counter for requests.",
		Labels: []string{"method", "endpoint"},
	})

	if err := chassis.Run(); err != nil {
		panic(err)
	}
}
