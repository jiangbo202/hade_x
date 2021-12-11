/**
 * @Author: jiangbo
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/11/26 10:31 下午
 */

package orm

import (
	"context"
	"github.com/jiangbo202/hade_x/framework"
	"github.com/jiangbo202/hade_x/framework/contract"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

type HadeGorm struct {
	container framework.Container // 服务容器

	configPath string
	dbs        map[string]*gorm.DB // 容器服务
	gormConfig *gorm.Config        // gorm的配置文件，可以修改

	lock *sync.RWMutex
}

func NewGorm(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	dbs := make(map[string]*gorm.DB)
	lock := &sync.RWMutex{}
	return &HadeGorm{
		container: container,
		dbs:       dbs,
		lock:      lock,
	}, nil
}

func (app *HadeGorm) GetDB(option ...contract.DBOption) (*gorm.DB, error) {
	logger := app.container.MustMake(contract.LogKey).(contract.Log)

	// 读取默认配置
	config := GetBaseConfig(app.container)

	logService := app.container.MustMake(contract.LogKey).(contract.Log)

	// 设置Logger
	ormLogger := NewOrmLogger(logService)
	config.Config = &gorm.Config{
		Logger: ormLogger,
	}

	// option对opt进行修改
	for _, opt := range option {
		if err := opt(app.container, config); err != nil {
			return nil, err
		}
	}
	if config.Dsn == "" {
		dsn, err := config.FormatDsn()
		if err != nil {
			return nil, err
		}
		config.Dsn = dsn
	}
	// fmt.Println(config.Dsn)  // ?allowNativePasswords=true

	if db, ok := app.dbs[config.Dsn]; ok {
		return db, nil
	}

	var db *gorm.DB
	var err error
	switch config.Driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.Dsn), config)
	case "postgres":
		db, err = gorm.Open(postgres.Open(config.Dsn), config)
		// case "sqlite":
		// 	db, err = gorm.Open(sqlite.Open(config.Dsn), gormConfig)
		// case "sqlserver":
		// 	db, err = gorm.Open(sqlserver.Open(config.Dsn), gormConfig)
		// case "clickhouse":
		// 	db, err = gorm.Open(clickhouse.Open(config.Dsn), gormConfig)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}

	if config.ConnMaxIdle > 0 {
		sqlDB.SetMaxIdleConns(config.ConnMaxIdle)
	}
	if config.ConnMaxOpen > 0 {
		sqlDB.SetMaxOpenConns(config.ConnMaxOpen)
	}
	if config.ConnMaxLifetime != "" {
		liftTime, err := time.ParseDuration(config.ConnMaxLifetime)
		if err != nil {
			logger.Error(context.Background(), "conn max lift time error", map[string]interface{}{
				"err": err,
			})
		} else {
			sqlDB.SetConnMaxLifetime(liftTime)
		}
	}

	if err != nil {
		app.dbs[config.Dsn] = db
	}

	return db, err
}
