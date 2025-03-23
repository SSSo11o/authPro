package handlers

import (
	"auth/internal/config"
	"auth/internal/model"
	"auth/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context) {
	var user model.User

	log.Println("Начало обработки запроса для создания")

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Ошибка при привязке JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать"})
		return
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Ошибка при преобразования")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при преобразования"})
		return
	}

	user.Password = hashPassword
	user.UpdatedAt = time.Now()

	result := config.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Ошибка при создании user: %v", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Успешно создан user: %+v", user)
	c.JSON(http.StatusOK, user)
}
