package templates

var GlobalFileTempalte = `package global

import (
	"{{.pkgname}}/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)


var (
	ConfigPath = "./config/config.yaml"
	GlobalC *config.Config
	DB *gorm.DB
	RDB *redis.Client
)
`
