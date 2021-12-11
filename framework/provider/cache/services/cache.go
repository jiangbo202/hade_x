/**
 * @Author: jiangbo
 * @Description:
 * @File:  cache
 * @Version: 1.0.0
 * @Date: 2021/12/05 12:58 下午
 */

package services

import (
    "github.com/pkg/errors"
    "time"
)

const (
    NoneDuration = time.Duration(-1)
)

var ErrKeyNotFound = errors.New("key not found")
var ErrTypeNotOk = errors.New("val type not ok")
