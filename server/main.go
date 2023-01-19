package main 

import (
"fmt"
 "github.com/gofiber/fiber/v2"
 "log"
)

type Itin struct {
	ID			int		`json:"id"`
	CarPrice 	int		`json:"car"`
	CarType		string	`json:"carType"`
	PlanePrice	int		`json:"planePrice"`
	PlaneType	string	`json:"planeType"`
	HotelPrice	int		`json:"HotelPrice"`
	HotelType	string	`json:"HotelType"`
}

func main (){

	app := fiber.New()

	itins := []Itin{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error{
		return c.SendString("ok")
	})

	app.Post("/api/itin", func(c *fiber.Ctx) error {
		itin := &Itin{}

		if err := c.BodyParser(itin); err != nil {
			return err
		}

		itin.ID = len(itins) + 1
		itins = append(itins, *itin)

		return c.JSON(itins)
	})

	log.Fatal(app.Listen(":4000"))
}
