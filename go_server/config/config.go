package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var Cfg = &Config{}

type Config struct {
	GormDB *gorm.DB
}

func init() {
	godotenv.Load()
}
