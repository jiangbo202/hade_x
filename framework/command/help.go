/**
 * @Author: jiangbo
 * @Description:
 * @File:  help
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:55 下午
 */

package command

import (
	"fmt"
	"github.com/gojiangbo/jiangbo/framework/cobra"
	"github.com/gojiangbo/jiangbo/framework/contract"
)

// helpCommand show current envionment
var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo for framework",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder:", appService.BaseFolder())
	},
}
