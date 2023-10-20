package global

import (
	"gorm.io/gorm"
	"grs-server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
