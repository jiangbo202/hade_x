package framework

/**
 * @Author: jiangbo
 * @Description:
 * @File:
 * @Version: 1.0.0
 * @Date: 2021/10/28 下午
 */

type NewInstance func(...interface{}) (interface{}, error)

// 服务提供者接口定义
type ServiceProvider interface {
	Name() string
	Register(Container Container) NewInstance
	Params(Container Container) []interface{}
	IsDefer() bool // false表示不延迟实例化
	Boot(Container Container) error
}
