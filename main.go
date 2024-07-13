package main

import (
	"go-fiber-sql/src/configuration"
	ds "go-fiber-sql/src/domain/datasources"
	repo "go-fiber-sql/src/domain/repositories"
	"go-fiber-sql/src/gateways"
	"go-fiber-sql/src/middlewares"
	sv "go-fiber-sql/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/watchakorn-18k/scalar-go"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)

	app.Use("/api/users/docs", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.yml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "bn-crud-admin API",
			},
			Theme:    "purple",
			Layout:   "classic",
			DarkMode: true,
		})
		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(htmlContent)
	})
	app.Use(recover.New())
	app.Use(middlewares.NewCORSMiddleware())

	sqlDb := ds.NewSqlDb()
	defer sqlDb.Connect.Close()
	userRepo := repo.NewUsersRepository(sqlDb)

	sv0 := sv.NewUsersService(userRepo)

	gateways.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
