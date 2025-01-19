package app_sql_query

import (
	"context"
	"encoding/json"
	"runtime/debug"

	pb "github.com/bianhuOK/api_client/app/app_sql_query/proto" // 导入生成的proto
	"github.com/bianhuOK/api_client/internal/domain/iface"
	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/bianhuOK/api_client/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcSqlController 实现 proto 中定义的 SqlQueryService 服务
type GrpcSqlController struct {
	pb.UnimplementedSqlQueryServiceServer // 嵌入生成的基础结构
	SqlTemplateService                    domain.TemplateService
	SqlQueryService                       iface.SqlQueryServiceIface
}

// NewGrpcSqlController 创建 gRPC 控制器实例
func NewGrpcSqlController(sqlTemplateSrv domain.TemplateService, sqlQuerySrv iface.SqlQueryServiceIface) *GrpcSqlController {
	return &GrpcSqlController{
		SqlTemplateService: sqlTemplateSrv,
		SqlQueryService:    sqlQuerySrv,
	}
}

// ExecuteQuery 实现 proto 中定义的 rpc 方法
func (c *GrpcSqlController) ExecuteQuery(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	logger := utils.GetLogger()

	// 延迟执行的函数，用于捕获并处理panic
	defer func() {
		if err := recover(); err != nil {
			logger.WithFields(map[string]interface{}{
				"panic": err,
				"stack": string(debug.Stack()),
			}).Error("handle grpc request panic")
		}
	}()

	// 解析请求参数
	var params map[string]interface{}
	if err := json.Unmarshal(req.Params, &params); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid params format")
	}

	logger.Info("ExecuteQuery", "app_id: ", req.AppId, ", params: ", params)

	// 根据app_id和params获取SQL模板
	template, err := c.SqlTemplateService.GetSqlTemplate(req.AppId, params)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Template not found")
	}

	// 执行SQL查询
	results, err := c.SqlQueryService.ExecuteSql(template.DbConfig, string(template.TemplateContent))
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to execute query")
	}

	// 序列化查询结果
	resultBytes, err := json.Marshal(results)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to serialize results")
	}

	return &pb.QueryResponse{
		Result: resultBytes,
	}, nil
}
