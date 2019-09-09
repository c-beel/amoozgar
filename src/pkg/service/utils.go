package service

import (
	"github.com/jinzhu/gorm"
	"github.com/c-beel/amoozgar/src/configman"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/validations"
)

func NewAmoozgarServer(cfg *configman.MainConfig) (*AmoozgarServer, error) {
	dbUri := cfg.DBConfig.GetDBUri()
	db, err := gorm.Open(cfg.DBConfig.Type, dbUri)
	validations.RegisterCallbacks(db)
	if err != nil {
		return nil, err
	}
	return &AmoozgarServer{
		DB: db,
	}, nil
}

func (server *AmoozgarServer) AutoMigrate() (err error) {
	return nil
}
