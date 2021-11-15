/**
 * @Author: jiangbo
 * @Description:
 * @File:  mapper
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:17 下午
 */

package demo

import (
	demoService "github.com/gojiangbo/jiangbo/app/provider/demo"
)

func UserModelsToUserDTOs(models []UserModel) []UserDTO {
	ret := []UserDTO{}
	for _, model := range models {
		t := UserDTO{
			ID:   model.UserId,
			Name: model.Name,
		}
		ret = append(ret, t)
	}
	return ret
}

func StudentsToUserDTOs(students []demoService.Student) []UserDTO {
	ret := []UserDTO{}
	for _, student := range students {
		t := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
