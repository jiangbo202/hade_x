/**
 * @Author: jiangbo
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:55 下午
 */

package command

import "github.com/gojiangbo/jiangbo/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app 命令
	root.AddCommand(initAppCommand())
	// env 命令
	root.AddCommand(initEnvCommand())
	// cron 命令
	root.AddCommand(initCronCommand())
	// config 命令
	root.AddCommand(initConfigCommand())
  // 编译命令
	root.AddCommand(initBuildCommand())
	// 调试命令
	root.AddCommand(initDevCommand())
	// 服务提供命令
  root.AddCommand(initProviderCommand())
  // 自动化命令
  root.AddCommand(initCmdCommand())
  // 自动化中间件迁移工具
  root.AddCommand(initMiddlewareCommand())
}
