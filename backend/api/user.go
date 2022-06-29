package api

import (
	"github.com/gin-gonic/gin"
	"monorepo-gin-typescript-react-vite/backend/service"
)

func GetAllUser(c *gin.Context) {
	users := service.GetAllUser()
	c.JSON(200, gin.H{"users": users})
}
