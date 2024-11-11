package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-repository/pkg/etcd"

	"go-zero-repository/user/api/internal/config"
	"go-zero-repository/user/api/internal/handler"
	"go-zero-repository/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ee := etcd.NewEtcd[config.TestConfig](c.Etcd)
	var err error
	c.TestConfig, err = ee.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.TestConfig.Name)
	ee.Listener(func() {
		fmt.Println(1111)
		c.TestConfig, _ = ee.GetConfig()
	})

	// 关闭日志
	logx.Disable()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
