package main

import (
	"fmt"
	"monorepo-gin-typescript-react-vite/backend/database"
	"monorepo-gin-typescript-react-vite/backend/service"
	"monorepo-gin-typescript-react-vite/backend/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func main() {
	app := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	app.Use(cors.New(corsConfig))
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	database_host := viper.GetString("backend.database_host")
	database_port := viper.GetString("backend.database_port")
	database_user := viper.GetString("backend.database_user")
	database_password := viper.GetString("backend.database_password")
	database_name := viper.GetString("backend.database_name")

	db := database.Connect(&database.DbConfig{
		Host: database_host, Port: database_port, User: database_user, Password: database_password, Name: database_name,
	})

	database.Migrate()

	// init service
	service.InitUserService(db)

	// api
	apiRouter := app.Group("/api")
	{
		apiRouter.GET("/users", api.GetAllUser)
	}

	baseUrl := viper.GetString("backend.baseUrl")
	port := viper.GetString("backend.port")

	app.Run(fmt.Sprintf("%s:%s", baseUrl, port))
}
