package xxx1

import (
	"fmt"

	"github.com/jiangbo202/hade_x/framework/cobra"
)

var Xxx1Command = &cobra.Command{
	Use:   "xxx1",
	Short: "xxx1",
	RunE: func(c *cobra.Command, args []string) error {
        container := c.GetContainer()
		fmt.Println(container)
		return nil
	},
}

