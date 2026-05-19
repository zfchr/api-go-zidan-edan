package main

import (
	"api-zidan-edan/api/config"
	"api-zidan-edan/api/models"
	"api-zidan-edan/api/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("SECRET", "MY_SECRET")

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.Routes(r)

	r.Run(":8080")
}
