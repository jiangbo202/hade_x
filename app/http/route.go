/**
 * @Author: jiangbo
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:15 下午
 */

package http

import (
	"github.com/gojiangbo/jiangbo/app/http/module/demo"
	"github.com/gojiangbo/jiangbo/framework/gin"
	"github.com/gojiangbo/jiangbo/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	demo.Register(r)
}