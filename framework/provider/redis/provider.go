/**
 * @Author: jiangbo
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/12/05 12:14 下午
 */

package redis

import (
    "github.com/jiangbo202/hade_x/framework"
    "github.com/jiangbo202/hade_x/framework/contract"
)

// RedisProvider 提供App的具体实现方法
type RedisProvider struct {
}

// Register 注册方法
func (h *RedisProvider) Register(container framework.Container) framework.NewInstance {
    return NewHadeRedis
}

// Boot 启动调用
func (h *RedisProvider) Boot(container framework.Container) error {
    return nil
}

// IsDefer 是否延迟初始化
func (h *RedisProvider) IsDefer() bool {
    return true
}

// Params 获取初始化参数
func (h *RedisProvider) Params(container framework.Container) []interface{} {
    return []interface{}{container}
}

// Name 获取字符串凭证
func (h *RedisProvider) Name() string {
    return contract.RedisKey
}



