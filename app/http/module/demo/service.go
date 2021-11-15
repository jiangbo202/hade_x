/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:18 下午
 */

package demo

type Service struct {
	repository *Repository
}


func NewService() *Service {
	repository := NewRepository()
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetUsers() []UserModel {
	ids := s.repository.GetUserIds()
	return s.repository.GetUserByIds(ids)
}