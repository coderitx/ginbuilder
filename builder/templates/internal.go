package templates

var InitMysqlTemplate = `package internal

import (
	"{{.pkgname}}/config/internal_config"
	"{{.pkgname}}/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitDB 初始化database
func InitDB(c internal_config.MysqlConfig) {
	var mysqlLogLeve logger.LogLevel
	if c.LogLevel == "info" {
		mysqlLogLeve = logger.Info
	} else {
		mysqlLogLeve = logger.Error
	}
	logx := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      mysqlLogLeve,
			Colorful:      true,
		})
	dbCfg := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DB),
	}
	db, err := gorm.Open(mysql.New(dbCfg), &gorm.Config{Logger: logx})
	if err != nil {
		zap.S().Errorf("connect mysql error: %v", err)
		return
	}
	db.Model(true)
	global.DB = db
	return
}`

var InitRedisTemplate = `package internal

import (
	"{{.pkgname}}/config/internal_config"
	"{{.pkgname}}/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func InitRedis(c internal_config.RedisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        c.Host,
		Password:    c.Password,
		DialTimeout: time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout: time.Duration(c.ReadTimeout) * time.Second,
		PoolSize:    c.PoolSize,
		PoolTimeout: time.Duration(c.PoolTimeout) * time.Second,
		MaxConnAge:  time.Duration(c.MaxConnAge) * time.Second,
	})

	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		zap.S().Errorf("redis connection error: %v", err)
		return
	}
	fmt.Println("=====redis========")
	global.RDB = rdb
	return
}`
