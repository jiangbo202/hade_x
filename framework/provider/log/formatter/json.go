/**
 * @Author: jiangbo
 * @Description:
 * @File:  json
 * @Version: 1.0.0
 * @Date: 2021/10/31 8:16 下午
 */

package formatter

import (
	"bytes"
	"encoding/json"
	"github.com/gojiangbo/jiangbo/framework/contract"
	"github.com/pkg/errors"
	"time"
)

func JsonFormatter(level contract.LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	fields["msg"] = msg
	fields["level"] = level
	fields["timestamp"] = t.Format(time.RFC3339)
	c, err := json.Marshal(fields)
	if err != nil {
		return bf.Bytes(), errors.Wrap(err, "json format error")
	}

	bf.Write(c)
	return bf.Bytes(), nil
}