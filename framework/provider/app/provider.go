/**
 * @Author: jiangbo
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/10/30 11:38 上午
 */

package app

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
)

type JiangAppProvider struct {
	BaseFolder string
}

func (j JiangAppProvider) Name() string {
	return contract.AppKey
}

func (j JiangAppProvider) Register(c framework.Container) framework.NewInstance {
	return NewJiangApp
}

func (j JiangAppProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c, j.BaseFolder}
}

func (j JiangAppProvider) IsDefer() bool {
	return false
}

func (j JiangAppProvider) Boot(c framework.Container) error {
	return nil
}


