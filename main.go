package main

import (
	"github.com/gojiangbo/jiangbo/app/console"
	"github.com/gojiangbo/jiangbo/app/http"
	"github.com/gojiangbo/jiangbo/framework"
	"github.com/gojiangbo/jiangbo/framework/provider/app"
	"github.com/gojiangbo/jiangbo/framework/provider/config"
	"github.com/gojiangbo/jiangbo/framework/provider/distributed"
	"github.com/gojiangbo/jiangbo/framework/provider/env"
	"github.com/gojiangbo/jiangbo/framework/provider/kernel"
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
