/**
 * @Author: jiangbo
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:15 下午
 */

package http

import (
  "github.com/jiangbo202/hade_x/app/http/middleware/cors"
  "github.com/jiangbo202/hade_x/app/http/module/demo"
  "github.com/jiangbo202/hade_x/framework/contract"
  "github.com/jiangbo202/hade_x/framework/gin"
  ginSwagger "github.com/jiangbo202/hade_x/framework/middleware/gin-swagger"
  "github.com/jiangbo202/hade_x/framework/middleware/gin-swagger/swaggerFiles"
  "github.com/jiangbo202/hade_x/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

  container := r.GetContainer()
  configService := container.MustMake(contract.ConfigKey).(contract.Config)

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
  r.Use(cors.Default())

  // 如果配置了swagger，则显示swagger的中间件
  if configService.GetBool("app.swagger") == true {
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  }
	demo.Register(r)
}
