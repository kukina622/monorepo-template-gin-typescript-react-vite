package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"monorepo-gin-typescript-react-vite/backend/model"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var repository *gorm.DB

func Connect(DbConfig *DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbConfig.User,
		DbConfig.Password,
		"tcp",
		DbConfig.Host,
		DbConfig.Port,
		DbConfig.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repository = db
	return db
}

func Migrate() {
	repository.AutoMigrate(&model.Users{})
}
