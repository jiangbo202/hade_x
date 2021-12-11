/**
 * @Author: jiangbo
 * @Description:
 * @File:  model
 * @Version: 1.0.0
 * @Date: 2021/10/30 2:18 下午
 */

package demo

import (
	"database/sql"
	"time"
)

type UserModel struct {
	UserId int
	Name   string
	Age    int
}

// User is gorm model
type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
