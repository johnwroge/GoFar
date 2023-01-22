// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// )

// type Itin struct {
// 	ID			int		`json:"id"`
// 	CarPrice 	int		`json:"car"`
// 	CarType		string	`json:"carType"`
// 	PlanePrice	int		`json:"planePrice"`
// 	PlaneType	string	`json:"planeType"`
// 	HotelPrice	int		`json:"HotelPrice"`
// 	HotelType	string	`json:"HotelType"`
// }

// func main (){

// 	app := fiber.New()

// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "http://localhost:3000",
// 		AllowHeaders: "Origin, Content-Type, Accept",
// 	}))

// 	itins := []Itin{}

// 	app.Get("/healthcheck", func(c *fiber.Ctx) error{
// 		return c.SendString("ok")
// 	})

// 	app.Post("/api/itin", func(c *fiber.Ctx) error {
// 		itin := &Itin{}

// 		if err := c.BodyParser(itin); err != nil {
// 			return err
// 		}

// 		itin.ID = len(itins) + 1
// 		itins = append(itins, *itin)

// 		return c.JSON(itins)
// 	})


// // need to change endpoints for todos to itins 
// 	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
// 		id, err := c.ParamsInt((id))

// 		if err != nil {
// 			return c.Status(401).SendString(("Invalid id"))
// 		}

// 		for i, t := range itins {
// 			if t.ID == id {
// 				itins[i].Done = true
// 				break
// 			}
// 		}
// 		return c.JSON(todos)
// 	})

// 	app.Get("/api/todos", func(c *fiber.Ctx) error {
// 		return c.JSON(todos)
// 	})



// 	log.Fatal(app.Listen(":4000"))
// }


package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Hello world")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)

	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))

}