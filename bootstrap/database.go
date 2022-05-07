package bootstrap

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hanya-go/app/models/user"
	"hanya-go/pkg/config"
	"hanya-go/pkg/database"
	"hanya-go/pkg/logger"
	"time"
)

// SetupDB 初始化数据库和ORM
func SetupDB() {
	var dbConfig gorm.Dialector

	switch config.Get("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)

		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("database connect not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 数据库结构迁移
	database.DB.AutoMigrate(&user.User{})
}
