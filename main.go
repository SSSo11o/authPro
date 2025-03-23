package main

import (
	"auth/internal/config"
	"auth/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config.ConnectDB()

	r := gin.Default()
	r.POST("/login", handlers.CreateUser)

	log.Println("Сервер запущен на порту 8080...")
	r.Run(":8080")
}
