/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/10/30 1:49 下午
 */

package kernel

import (
	"github.com/jiangbo202/hade_x/framework/gin"
	"net/http"
)

// JiangKernelService 引擎服务
type JiangKernelService struct {
	engine *gin.Engine
}

func (j JiangKernelService) HttpEngine() http.Handler {
	return j.engine
}

func NewJiangKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &JiangKernelService{
		engine: httpEngine,
	}, nil
}

