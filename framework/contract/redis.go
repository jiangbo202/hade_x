/**
 * @Author: jiangbo
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2021/12/05 11:09 上午
 */

package contract

import (
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/jiangbo202/hade_x/framework"
)

const RedisKey = "hade:redis"

type RedisOption func(container framework.Container, config * RedisConfig) error

type RedisService interface {
    GetClient(option ...RedisOption) (*redis.Client, error)
}

type RedisConfig struct {
    *redis.Options
}

// UniqKey 用来唯一标识一个RedisConfig配置
func (config *RedisConfig) UniqKey() string {
    return fmt.Sprintf("%v_%v_%v_%v", config.Addr, config.DB, config.Username, config.Network)
}