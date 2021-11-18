/**
 * @Author: jiangbo
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/10/30 1:46 下午
 */

package kernel

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
	"github.com/jiangbo202/hade_x/framework/gin"
)

type JiangKernelProvider struct {
	HttpEngine *gin.Engine
}

func (j JiangKernelProvider) Name() string {
	return contract.KernelKey
}

func (j JiangKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewJiangKernelService
}

func (j JiangKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{j.HttpEngine}
}

func (j JiangKernelProvider) IsDefer() bool {
	return false
}

func (j JiangKernelProvider) Boot(c framework.Container) error {
	if j.HttpEngine == nil {
		j.HttpEngine = gin.Default()
	}
	j.HttpEngine.SetContainer(c)
	return nil
}



