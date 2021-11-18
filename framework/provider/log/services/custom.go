/**
 * @Author: jiangbo
 * @Description:
 * @File:  custom
 * @Version: 1.0.0
 * @Date: 2021/10/31 8:38 下午
 */

package services

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
	"io"
)

type HadeCustomLog struct {
	HadeLog
}

func NewHadeCustomLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &HadeConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c
	return log, nil
}
