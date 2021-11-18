/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/10/30 11:49 上午
 */

package app

import (

	"flag"
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/util"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"path/filepath"
)

type JiangApp struct {
	c          framework.Container
	baseFolder string
	jId      string // 表示当前这个j的唯一id, 可以用于分布式锁等
	configMap map[string]string // 配置加载
}

// AppID 表示这个App的唯一ID
func (j JiangApp) AppID() string {
	return j.jId
}

func (j JiangApp) Version() string {
	return "0.0.0"
}

func (j JiangApp) BaseFolder() string {
	if j.baseFolder != "" {
		return j.baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (j JiangApp) ConfigFolder() string {
	if val, ok := j.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(j.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (j JiangApp) LogFolder() string {
	if val, ok := j.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(j.StorageFolder(), "log")
}

func (j JiangApp) HttpFolder() string {
	if val, ok := j.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(j.BaseFolder(), "app", "http")
}

func (j JiangApp) ConsoleFolder() string {
	if val, ok := j.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(j.BaseFolder(), "app", "console")
}

func (j JiangApp) StorageFolder() string {
	if val, ok := j.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(j.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (j JiangApp) ProviderFolder() string {
	if val, ok := j.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(j.BaseFolder(), "app", "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (j JiangApp) MiddlewareFolder() string {
	if val, ok := j.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(j.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (j JiangApp) CommandFolder() string {
	if val, ok := j.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(j.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (j JiangApp) RuntimeFolder() string {
	if val, ok := j.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(j.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (j JiangApp) TestFolder() string {
	if val, ok := j.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(j.BaseFolder(), "test")
}

func NewJiangApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param count not OK")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	// 如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
		flag.Parse()
	}
	jId := uuid.New().String()
	configMap := map[string]string{}
	return &JiangApp{baseFolder: baseFolder, c: container, jId: jId, configMap: configMap}, nil
}

func (j JiangApp) LoadAppConfig(kv map[string]string) {
	j.configMap = kv
}

// AppFolder 代表app目录
func (j *JiangApp) AppFolder() string {
  if val, ok := j.configMap["app_folder"]; ok {
    return val
  }
  return filepath.Join(j.BaseFolder(), "app")
}
