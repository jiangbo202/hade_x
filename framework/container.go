package framework

import (
	"errors"
	"fmt"
	"sync"
)

/**
 * @Author: iiangbo
 * @Description:
 * @File:
 * @Version: 1.0.0
 * @Date: 2021/10/28 下午
 */

type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，不返回 error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会 panic。
	//所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	//它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	//这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// JContainer 服务容器的具体实现
type JContainer struct {
	Container // 强制要求  实现 Container 接口
	// providers 存储注册的服务提供者，key 为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key 为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewInsContainer() *JContainer {
	return &JContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock: sync.RWMutex{},
	}
}

func (i *JContainer) PrintProviders() []string {
	var ret []string
	for _, provider := range i.providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

func (i *JContainer) NameList() []string {
  var ret []string
  for _, provider := range i.providers {
    name := provider.Name()
    ret = append(ret, name)
  }
  return ret
}

func (i *JContainer) Bind(provider ServiceProvider) error {
	i.lock.Lock()
	key := provider.Name()

	i.providers[key] = provider
	i.lock.Unlock()

	if provider.IsDefer() == false {
		if err := provider.Boot(i); err != nil {
			return err
		}
		// 实例化
		params := provider.Params(i)
		method := provider.Register(i)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		i.instances[key] = instance
	}
	return nil
}

func (i *JContainer) IsBind(key string) bool {
	return i.findServiceProvider(key) != nil
}

func (i *JContainer) findServiceProvider(key string) ServiceProvider {
	i.lock.RLock()
	defer i.lock.RUnlock()
	if sp, ok := i.providers[key]; ok {
		return sp
	}
	return nil
}

func (i *JContainer) MustMake(key string) interface{} {
	serv, err := i.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (i *JContainer) Make(key string) (interface{}, error) {
	return i.make(key, nil, false)
}

// MakeNew 方式使用内部的 make 初始化
func (i *JContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return i.make(key, params, true)
}


func (i *JContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(i); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(i)
	}
	method := sp.Register(i)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

func (i *JContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := i.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return i.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := i.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := i.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	i.instances[key] = inst
	return inst, nil
}

