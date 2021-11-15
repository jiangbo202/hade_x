/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:20 下午
 */

package demo

import "github.com/gojiangbo/jiangbo/framework"

type Service struct {
	container framework.Container
}

func NewService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &Service{container: container}, nil
}

func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
