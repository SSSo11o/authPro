package config

import (
	"auth/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("База данных успешно подключена")
			if err := DB.AutoMigrate(&model.Person{}); err != nil {
				log.Fatal("Ошибка при миграции базы данных: ", err)
			}
			log.Println("Миграция базы данных успешно выполнена")
			return
		}
		log.Printf("Ошибка подключения к базе данных: %v, попытка %d\n", err, i+1)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Не удалось подключиться к базе данных после нескольких попыток")
}
