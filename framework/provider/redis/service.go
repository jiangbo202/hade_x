/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/12/05 12:14 下午
 */

package redis

import (
    "github.com/go-redis/redis/v8"
    "github.com/jiangbo202/hade_x/framework"
    "github.com/jiangbo202/hade_x/framework/contract"
    "sync"
)

type HadeRedis struct {
    container framework.Container
    clients map[string]*redis.Client
    lock *sync.RWMutex
}

func NewHadeRedis(params ...interface{}) (interface{}, error) {
    container := params[0].(framework.Container)
    clients := make(map[string]*redis.Client)
    lock := &sync.RWMutex{}
    return &HadeRedis{container: container,
        clients: clients, lock: lock,
        }, nil
}


func (h *HadeRedis) GetClient(option ...contract.RedisOption) (*redis.Client, error) {
    config := GetBaseConfig(h.container)

    for _, opt := range option {
        if err := opt(h.container, config); err !=nil{
            return nil, err
        }
    }

    key := config.UniqKey()

    h.lock.RLock()
    if db, ok := h.clients[key]; ok {
        h.lock.RUnlock()
        return db, nil
    }
    h.lock.RUnlock()
    h.lock.Lock()
    defer h.lock.Unlock()

    client := redis.NewClient(config.Options)
    h.clients[key] = client
    return client, nil
}




