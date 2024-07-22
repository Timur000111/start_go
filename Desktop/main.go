package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Определение структуры Message
type Message struct {
	Id      uint   `json:"id"`
	Person  string `json:"person"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// Функция, возвращающая текущее время в формате строки
func realTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

// Функция для поиска человека по ID в map
func searchPerson(persons map[uint]string, needID uint) (uint, string) {
	for id, person := range persons {
		if id == needID {
			return id, person
		}
	}
	return 0, ""
}

func main() {
	persons := map[uint]string{
		1234567890: "Tim",
		1234567891: "Liz",
		1234567892: "PETR",
		1234567893: "MASH",
		1234567894: "Kolz",
	}

	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		// Получение текущего времени
		currentTime := realTime()

		// Поиск человека по ID
		id, person := searchPerson(persons, 1234567890)

		// Создание экземпляра структуры Message
		msg := Message{
			Id:      id,
			Person:  person,
			Message: "Hello World!",
			Time:    currentTime,
		}

		// Отправка JSON-ответа с использованием структуры Message
		context.JSON(http.StatusOK, msg)
	})

	// Запуск сервера на порту 8080
	route.Run(":8081")
}
