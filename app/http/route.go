/**
 * @Author: jiangbo
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:15 下午
 */

package http

import (
	"github.com/jiangbo202/hade_x/app/http/module/demo"
	"github.com/jiangbo202/hade_x/framework/gin"
	"github.com/jiangbo202/hade_x/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	demo.Register(r)
}
