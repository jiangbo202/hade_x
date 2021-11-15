/**
 * @Author: jiangbo
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/30 1:43 下午
 */

package contract

import (
	"net/http"
)

const KernelKey = "jiang:kernel"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	// HttpEngine 提供gin的Engine结构
	HttpEngine() http.Handler
}
