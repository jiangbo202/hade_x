/**
 * @Author: jiangbo
 * @Description:
 * @File:  prefix
 * @Version: 1.0.0
 * @Date: 2021/10/31 8:17 下午
 */

package formatter

import "github.com/jiangbo202/hade_x/framework/contract"

func Prefix(level contract.LogLevel) string {
	prefix := ""
	switch level {
	case contract.PanicLevel:
		prefix = "[Panic]"
	case contract.FatalLevel:
		prefix = "[Fatal]"
	case contract.ErrorLevel:
		prefix = "[Error]"
	case contract.WarnLevel:
		prefix = "[Warn]"
	case contract.InfoLevel:
		prefix = "[Info]"
	case contract.DebugLevel:
		prefix = "[Debug]"
	case contract.TraceLevel:
		prefix = "[Trace]"
	}
	return prefix
}
