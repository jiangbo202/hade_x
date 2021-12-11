/**
 * @Author: jiangbo
 * @Description:
 * @File:  env
 * @Version: 1.0.0
 * @Date: 2021/10/31 2:24 下午
 */

package test

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/provider/app"
)

const (
	BasePath = "/Users/jiangbo/jiang/go_2021_dir/from0/gk02/"
)


func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewInsContainer()
	// 绑定App服务提供者
	container.Bind(&app.JiangAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	// container.Bind(&env.HadeTestingEnvProvider{})
	return container
}