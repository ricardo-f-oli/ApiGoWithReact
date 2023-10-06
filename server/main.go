package main

import (
	_ "fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Todo struct {
	ID    int `json:"id"`
	Title int `json:"title"`
	Done  int `json:"done"`
	Body  int `json:"body"`
}

func main() {
	app := fiber.New()

	app.Get("/heathcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	todos := []Todo{}

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(todos)

	})
	log.Fatal(app.Listen(":4000"))
}
