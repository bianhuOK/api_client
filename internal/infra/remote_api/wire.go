package remoteapi

import (
	"github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/google/wire"
)

func ProvideApiURL() string {
	// 从配置文件或环境变量获取
	return "http://your-api-url"
}

var RemoteApiSet = wire.NewSet(
	ProvideApiURL,
	NewSqlApiRest,
	wire.Bind(new(iface.RemoteAPI), new(*SqlApiRest)),
)

var MockRemoteApiSet = wire.NewSet(
	NewMockRemoteAPI,
	wire.Bind(new(iface.RemoteAPI), new(*MockRemoteAPI)),
)
