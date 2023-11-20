package main

import "github.com/gofiber/fiber/v2"

func main() {
	//обращаемся к фреймворку
	rog := fiber.New()

	// обрщаемся к переменной rog c методом запроса GET с микрофреймворком
	rog.Get("/", func(c *fiber.Ctx) error {
		// возвращаем строку Hello World
		return c.SendString("Hello world")
	})

	// Запуск сервера на 127.0.0.1:3000
	err := rog.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}
