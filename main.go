package main

import (
	"github.com/jiangbo202/hade_x/app/console"
	"github.com/jiangbo202/hade_x/app/http"
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/provider/app"
	"github.com/jiangbo202/hade_x/framework/provider/cache"
	"github.com/jiangbo202/hade_x/framework/provider/config"
	"github.com/jiangbo202/hade_x/framework/provider/distributed"
	"github.com/jiangbo202/hade_x/framework/provider/env"
	"github.com/jiangbo202/hade_x/framework/provider/kernel"
	"github.com/jiangbo202/hade_x/framework/provider/log"
	"github.com/jiangbo202/hade_x/framework/provider/orm"
	"github.com/jiangbo202/hade_x/framework/provider/redis"
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
	container.Bind(&log.JiangLogServiceProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.HadeCacheProvider{})
	container.Bind(&orm.GormProvider{})

	if engine, err := http.NewHttpEngine(container);err == nil {
		container.Bind(&kernel.JiangKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
}
