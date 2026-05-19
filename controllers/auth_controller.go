package controllers

import (
	"net/http"

	"api-zidan-edan/api/config"

	"api-zidan-edan/api/middleware"

	"api-zidan-edan/api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req models.User
	c.BindJSON(&req)

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	user := models.User{Name: req.Name, Email: req.Email, Password: string(hash)}

	config.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var req models.User
	c.BindJSON(&req)

	var user models.User
	config.DB.Where("email = ?", req.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := middleware.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
