package main

import (
	"log"

	"github.com/Anvarsha-k/restfulEcommerceFiber/database"
	"github.com/Anvarsha-k/restfulEcommerceFiber/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to golang")
}

func setupRoutes(app *fiber.App) {
	//user routes
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users/list", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/update/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//product Routes
	app.Post("/api/product/create",routes.CreateProduct)
	app.Get("/api/product/list",routes.GetProducts)
	app.Get("api/product/:id", routes.GetProduct)
	app.Put("api/product/update/:id",routes.UpdateProduct)
	app.Delete("api/product/delete/:id",routes.DeleteProduct)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
