/**
 * @Author: jiangbo
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/11/26 10:31 下午
 */

package orm

import (
    "github.com/jiangbo202/hade_x/framework"
    "github.com/jiangbo202/hade_x/framework/contract"
)

type GormProvider struct {
}

func (j GormProvider) Name() string {
    return contract.ORMKey
}

func (j GormProvider) Register(c framework.Container) framework.NewInstance {
    return NewGorm
}

func (j GormProvider) Params(c framework.Container) []interface{} {
    return []interface{}{c}
}

func (j GormProvider) IsDefer() bool {
    return false
}

func (j GormProvider) Boot(c framework.Container) error {
    return nil
}


