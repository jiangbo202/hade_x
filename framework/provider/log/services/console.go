/**
 * @Author: jiangbo
 * @Description:
 * @File:  console
 * @Version: 1.0.0
 * @Date: 2021/10/31 8:38 下午
 */

package services

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
	"os"
)

// HadeConsoleLog 代表控制台输出
type HadeConsoleLog struct {
	HadeLog
}

// NewHadeConsoleLog 实例化HadeConsoleLog
func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &HadeConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
