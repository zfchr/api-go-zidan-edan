package routes

import (
	"api-zidan-edan/api/controllers"
	"api-zidan-edan/api/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	authorized := router.Group("/tasks")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/", controllers.GetTasks)
		authorized.POST("/", controllers.CreateTask)
		authorized.PUT("/:id", controllers.UpdateTask)
		authorized.DELETE("/:id", controllers.DeleteTask)
	}
}
