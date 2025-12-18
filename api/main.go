package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"acupofcoffee/api/internal/config"
	"acupofcoffee/api/internal/handler"
	"acupofcoffee/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Railway 使用环境变量 PORT
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			c.Port = p
		}
	}

	// 确保绑定到 0.0.0.0
	c.Host = "0.0.0.0"

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
