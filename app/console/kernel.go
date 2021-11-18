/**
 * @Author: jiangbo
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:42 下午
 */

package console

import (
	"github.com/jiangbo202/hade_x/app/console/command/demo"
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/cobra"
	"github.com/jiangbo202/hade_x/framework/command"
)

// RunCommand  初始化根Command并运行
func RunCommand(container framework.Container) error {
	// 根Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "jiang",
		// 简短介绍
		Short: "jiang 命令",
		// 根命令的详细介绍
		Long: "jiang 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根Command设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	// 执行RootCommand
	return rootCmd.Execute()
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.FooCommand)

	// 每秒调用一次Foo命令
	//rootCmd.AddCronCommand("* * * * * *", demo.FooCommand)

	// 启动一个分布式任务调度，调度的服务名称为init_func_for_test，每个节点每5s调用一次Foo命令，抢占到了调度任务的节点将抢占锁持续挂载2s才释放
	//rootCmd.AddDistributedCronCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCommand, 2*time.Second)
}
