package data

import (
	v1 "review-b/api/review/v1"
	"review-b/internal/conf"

	"context"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewReviewSeviceClient, NewData, NewBusinessRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	reviewclient v1.ReviewClient
	log          *log.Logger
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		reviewclient: rc,
	}, cleanup, nil
}

//	func NewReviewSeviceClient(c *conf.Data) v1.ReviewClient {
//		conn, err := grpc.DialInsecure(
//			context.Background(),
//			grpc.WithEndpoint("127.0.0.1:9092"),
//		)
//		if err != nil {
//			panic(err)
//		}
//		return v1.NewReviewClient(conn)
//	}

// 服务发现的discovery
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	if conf == nil || conf.Nacos == nil {
		panic("conf.Registry or conf.Registry.Nacos is nil")
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Ip, conf.Nacos.Port),
	}

	//连接nacos的客户端
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//不用加健康检查，自动检测
	r := nacos.New(client)

	return r
}
func NewReviewSeviceClient(d registry.Discovery) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///review-service.grpc"),
		grpc.WithDiscovery(d),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(conn)
}
