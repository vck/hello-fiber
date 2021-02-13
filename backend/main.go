package main

import "github.com/gofiber/fiber/v2"

type Todo struct {
	ID int
	Title string
}


func main() {

	db := []Todo{}

    app := fiber.New()

    prevID := 1

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("todo API v.0.1")
    })

    app.Get("/todos", func(c *fiber.Ctx) error {
    	return c.JSON(db)
    })

    app.Post("/add", func(c *fiber.Ctx) error {
    	var todo Todo
    	todo.ID = prevID
    	todo.Title = "Hello world"
    	db = append(db, todo)

    	prevID += 1
    	return c.JSON(todo)

    })

    app.Listen(":3000")
}
