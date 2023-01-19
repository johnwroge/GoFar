package main 

import (
"fmt"
 "github.com/gofiber/fiber/v2"
 "log"
)

func main (){
	fmt.Print("hello world asshole")
	app := fiber.New()

	app.Get("/healthchecl", func(c *fiber.Ctx) error{
		return c.SendString("ok")
	})

	log.Fatal(app.Listen(":4000"))
}
