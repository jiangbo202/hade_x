/**
 * @Author: jiangbo
 * @Description:
 * @File:  j
 * @Version: 1.0.0
 * @Date: 2021/10/31 10:18 上午
 */

package env

import (
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
)

type JiangEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (j *JiangEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewJiangEnv
}

// Boot will called when the service instantiate
func (j *JiangEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	j.Folder = app.BaseFolder()

	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (j *JiangEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (j *JiangEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{j.Folder}
}

/// Name define the name for this service
func (j *JiangEnvProvider) Name() string {
	return contract.EnvKey
}
