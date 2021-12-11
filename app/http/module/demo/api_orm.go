/**
 * @Author: jiangbo
 * @Description:
 * @File:  api_orm
 * @Version: 1.0.0
 * @Date: 2021/11/26 10:50 下午
 */

package demo

import (
    "database/sql"
    "github.com/jiangbo202/hade_x/framework/contract"
    "github.com/jiangbo202/hade_x/framework/gin"
    "github.com/jiangbo202/hade_x/framework/provider/orm"
    "time"
)


func (api *DemoApi) DemoOrm(c *gin.Context) {
    logger := c.MustMakeLog()
    logger.Info(c, "request start", nil)
    gormService := c.MustMake(contract.ORMKey).(contract.ORMService)
    db, err := gormService.GetDB(orm.WithConfigPath("database.default"))
    if err != nil {
        logger.Error(c, err.Error(), nil)
        c.AbortWithError(50001, err)
        return
    }
    db.WithContext(c)

    err = db.AutoMigrate(&User{})
    if err != nil {
        c.AbortWithError(500, err)
        return
    }
    logger.Info(c, "migrate ok", nil)

    email := "foo@gmail.com"
    name := "foo"
    age := uint8(25)
    birthday := time.Date(2001, 1, 1, 1, 1, 1, 1, time.Local)
    user := &User{
        Name:         name,
        Email:        &email,
        Age:          age,
        Birthday:     &birthday,
        MemberNumber: sql.NullString{},
        ActivatedAt:  sql.NullTime{},
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
    err = db.Create(user).Error
    logger.Info(c, "insert user", map[string]interface{}{
        "id":  user.ID,
        "err": err,
    })

    user.Name = "bar"
    err = db.Save(user).Error
    logger.Info(c, "update user", map[string]interface{}{
        "err": err,
        "id":  user.ID,
    })

    queryUser := &User{ID: user.ID}

    err = db.First(queryUser).Error
    logger.Info(c, "query user", map[string]interface{}{
        "err":  err,
        "name": queryUser.Name,
    })

    logger.Info(c, "delete user", map[string]interface{}{
        "err": err,
        "id":  user.ID,
    })
    c.JSON(200, "ok")
}

