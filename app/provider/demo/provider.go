/**
 * @Author: jiangbo
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:19 下午
 */

package demo

import (
	"github.com/jiangbo202/hade_x/framework"
)

type DemoProvider struct {
	framework.ServiceProvider

	c framework.Container
}

func (sp *DemoProvider) Name() string {
	return DemoKey
}

func (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (sp *DemoProvider) IsDefer() bool {
	return false
}

func (sp *DemoProvider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}

func (sp *DemoProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}
