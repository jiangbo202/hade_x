/**
 * @Author: jiangbo
 * @Description:
 * @File:  custom_context_contract
 * @Version: 1.0.0
 * @Date: 2021/10/31 2:22 下午
 */

package gin

import "github.com/jiangbo202/hade_x/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Context) MustMakeApp() contract.App {
	return c.MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Context) MustMakeKernel() contract.Kernel {
	return c.MustMake(contract.KernelKey).(contract.Kernel)
}
