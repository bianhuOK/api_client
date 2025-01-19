package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	pb "github.com/bianhuOK/api_client/app/app_sql_query/proto" // 导入生成的proto
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/client"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/core/common"
	_ "github.com/go-chassis/go-chassis/v2/middleware/accesslog"
	"github.com/go-chassis/openlog"
)

func TestGrpcSqlQuery(t *testing.T) {
	//export CHASSIS_HOME=/{path}/{to}/grpc/client/
	t.Log("TestGrpcSqlQuery")
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// 设置 CHASSIS_HOME 环境变量
	homeDir := filepath.Join(pwd)
	homeDir = filepath.Join(pwd, "..", "..", "..")
	os.Setenv("CHASSIS_HOME", homeDir)

	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	//declare reply struct
	reply := &pb.QueryResponse{}
	//header will transport to target service
	ctx := common.NewContext(map[string]string{
		"X-User": "pete",
	})
	invoker := core.NewRPCInvoker()

	//Invoke with micro service name, schema ID and operation ID
	params := map[string]interface{}{
		"a": "123",
	}
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	req := pb.QueryRequest{
		AppId:  "Peter",
		Params: paramsBytes,
	}
	if err := invoker.Invoke(ctx, "ApiMartServer", "sqlquery.SqlQueryService", "ExecuteQuery",
		&req, reply, core.WithProtocol("grpc")); err != nil {
		openlog.Error("error" + err.Error())
	}
	openlog.Info(string(reply.Result))
}
