package global

import (
	"gorm.io/gorm"
	"grs-server/internal/conf"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)
