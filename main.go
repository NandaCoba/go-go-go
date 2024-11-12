package main

import (
	"belajar/controllers"
	"belajar/db"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	db.Koneksi()

	app.Get("/", func(c *fiber.Ctx) error {
		data, err := controllers.GetAll()

		if err != nil {
			fmt.Println("gagal get all users")
		}

		return c.Status(200).JSON(fiber.Map{
			"data": data,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		name := c.FormValue("Name")
		age := c.FormValue("Age")
		ageInt, _ := strconv.Atoi(age)

		if name == "" {
			fmt.Println("nama tidak boleh kosong")
		}
		if ageInt < 0 {
			fmt.Println("Usia tidak boleh kosong")
		}

		createUser, err := controllers.Create(name, ageInt)

		if err != nil {
			fmt.Println("failed create new user")
		}

		return c.Status(200).JSON(fiber.Map{
			"data": createUser,
		})
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		name := c.FormValue("Name")
		age := c.FormValue("Age")
		ageInt, _ := strconv.Atoi(age)
		idInt, _ := strconv.Atoi(id)

		if name == "" {
			fmt.Println("nama tidak boleh kosong")
		}
		if ageInt < 0 {
			fmt.Println("Usia tidak boleh kosong")
		}

		updateUser, err := controllers.Update(idInt, name, ageInt)

		if err != nil {
			fmt.Println("failed create new user")
		}

		return c.Status(200).JSON(fiber.Map{
			"data": updateUser,
		})
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, _ := strconv.Atoi(id)
		deleteUser, _ := controllers.Delete(idInt)
		return c.Status(200).JSON(fiber.Map{
			"data": deleteUser,
		})
	})

	app.Listen(":3000")
}
