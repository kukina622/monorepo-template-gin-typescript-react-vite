package service

import (
	"gorm.io/gorm"
	"monorepo-gin-typescript-react-vite/backend/model"
)


type userService struct {
	repository *gorm.DB
}

var userServiceInstance userService

func InitUserService(db *gorm.DB) {
	userServiceInstance = userService{repository: db}
}


func GetAllUser() []model.Users {
	var users []model.Users
	userServiceInstance.repository.Find(&users)
	return users
}