/**
 * @Author: jiangbo
 * @Description:
 * @File:  custom_engine
 * @Version: 1.0.0
 * @Date: 2021/10/30 3:06 下午
 */

package gin

import "github.com/jiangbo202/hade_x/framework"

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container

}

// engine实现container的绑定封装
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

func (engine *Engine) GetContainer() framework.Container {
  return engine.container
}
