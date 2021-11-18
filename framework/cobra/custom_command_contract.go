/**
 * @Author: jiangbo
 * @Description:
 * @File:  custom_command_contract
 * @Version: 1.0.0
 * @Date: 2021/10/31 2:19 下午
 */

package cobra

import "github.com/jiangbo202/hade_x/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Command) MustMakeApp() contract.App {
	return c.GetContainer().MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Command) MustMakeKernel() contract.Kernel {
	return c.GetContainer().MustMake(contract.KernelKey).(contract.Kernel)
}
