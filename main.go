package main

import (
	"github.com/jiangbo202/hade_x/app/console"
	"github.com/jiangbo202/hade_x/app/http"
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/provider/app"
	"github.com/jiangbo202/hade_x/framework/provider/config"
	"github.com/jiangbo202/hade_x/framework/provider/distributed"
	"github.com/jiangbo202/hade_x/framework/provider/env"
	"github.com/jiangbo202/hade_x/framework/provider/kernel"
)

/**
 * @Author: jiangbo
 * @Description:
 * @File:
 * @Version: 1.0.0
 * @Date: 2021/10/20 下午
 */

func main() {

	container := framework.NewInsContainer()
	container.Bind(&app.JiangAppProvider{})
	container.Bind(&env.JiangEnvProvider{})

	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.HadeConfigProvider{})

	if engine, err := http.NewHttpEngine();err == nil {
		container.Bind(&kernel.JiangKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
}
