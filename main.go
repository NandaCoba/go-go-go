package main

import (
	"belajar/controller"
	"belajar/db"
	"belajar/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	db.Koneksi()

	app.Get("/", func(c *fiber.Ctx) error {
		data, err := controller.GetAll()
		if utils.SendError(err) {
			c.Status(500).JSON(fiber.Map{
				"message": "Failed get all data",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"data": data,
		})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, _ := strconv.Atoi(id)
		data, err := controller.GetId(idInt)
		if utils.SendError(err) {
			c.Status(500).JSON(fiber.Map{
				"message": "Failed get data user id",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"data": data,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		nama := c.FormValue("nama")
		usia := c.FormValue("usia")
		usiaInt, _ := strconv.Atoi(usia)

		data, err := controller.Create(nama, usiaInt)

		if utils.SendError(err) {
			utils.JsonError(nil, 500, "Failed create new user")
		}
		return c.Status(200).JSON(fiber.Map{
			"data": data,
		})
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		nama := c.FormValue("nama")
		usia := c.FormValue("usia")

		// convert to int
		usiaInt, _ := strconv.Atoi(usia)
		idInt, _ := strconv.Atoi(id)

		data, err := controller.Update(idInt, nama, usiaInt)

		if utils.SendError(err) {
			utils.JsonError(nil, 500, "Failed update user")
		}
		return c.Status(200).JSON(fiber.Map{
			"data": data,
		})
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// convert to int
		idInt, _ := strconv.Atoi(id)

		data, err := controller.Delete(idInt)

		if utils.SendError(err) {
			utils.JsonError(nil, 500, "Failed delete user")
		}
		return c.Status(200).JSON(fiber.Map{
			"data": data,
		})
	})

	app.Listen(":3000")
}
