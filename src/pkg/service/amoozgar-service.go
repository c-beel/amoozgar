package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type AmoozgarServer struct {
	DB *gorm.DB
}
