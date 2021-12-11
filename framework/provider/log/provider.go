/**
 * @Author: jiangbo
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/10/31 8:03 下午
 */

package log

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
	"github.com/jiangbo202/hade_x/framework/provider/log/formatter"
	"github.com/jiangbo202/hade_x/framework/provider/log/services"
	"io"
	"strings"
)

type JiangLogServiceProvider struct {
	framework.ServiceProvider

	Driver string
	// 日志级别
	Level contract.LogLevel
	// 日志输出格式方法
	Formatter contract.Formatter
	// 日志context上下文信息获取函数
	CtxFielder contract.CtxFielder
	// 日志输出信息
	Output io.Writer
}

func (j JiangLogServiceProvider) Name() string {
	return contract.LogKey
}

// Register 注册一个服务实例
func (j *JiangLogServiceProvider) Register(c framework.Container) framework.NewInstance {
	if j.Driver == "" {
		tcs, err := c.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用console
			return services.NewHadeConsoleLog
		}
		cs := tcs.(contract.Config)
		j.Driver = strings.ToLower(cs.GetString("log.Driver"))
	}
	// 根据driver的配置项确定
	switch j.Driver {
	case "single":
		return services.NewHadeSingleLog
	case "rotate":
		return services.NewHadeRotateLog
	case "console":
		return services.NewHadeConsoleLog
	case "custom":
		return services.NewHadeCustomLog
	default:
		return services.NewHadeConsoleLog
	}
}

// Params 定义要传递给实例化方法的参数
func (j *JiangLogServiceProvider) Params(c framework.Container) []interface{} {
	// 获取configService
	configService := c.MustMake(contract.ConfigKey).(contract.Config)

	// 设置参数formatter
	if j.Formatter == nil {
		j.Formatter = formatter.TextFormatter
		if configService.IsExist("log.formatter") {
			v := configService.GetString("log.formatter")
			if v == "json" {
				j.Formatter = formatter.JsonFormatter
			} else if v == "text" {
				j.Formatter = formatter.TextFormatter
			}
		}
	}

	if j.Level == contract.UnknownLevel {
		j.Level = contract.InfoLevel
		if configService.IsExist("log.level") {
			j.Level = logLevel(configService.GetString("log.level"))
		}
	}

	// 定义5个参数
	return []interface{}{c, j.Level, j.CtxFielder, j.Formatter, j.Output}
}

func (j JiangLogServiceProvider) IsDefer() bool {
	return false
}

func (j JiangLogServiceProvider) Boot(Container framework.Container) error {
	return nil
}

// logLevel get level from string
func logLevel(config string) contract.LogLevel {
	switch strings.ToLower(config) {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	}
	return contract.UnknownLevel
}
