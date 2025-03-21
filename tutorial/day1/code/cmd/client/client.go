package main

import (
	"context"
	"fmt"
	"hello/conf"
	"hello/kitex_gen/echo"
	"hello/kitex_gen/echo/echoservice"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}

	// 为了传输元信息数据，必须声明底层协议
	cli, err := echoservice.NewClient(conf.GetConf().Kitex.Service,
		client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		panic(err)
	}

	req := &echo.Req{
		Mes: "hello, world",
	}

	// 自定义含有元信息的上下文，从客户端传给服务端
	ctx := context.Background()
	ctx = metainfo.WithPersistentValue(ctx, "CLIENT_NAME", "echo_client")
	resp, err := cli.Echo(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
