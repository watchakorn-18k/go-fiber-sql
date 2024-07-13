package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCORSMiddleware() fiber.Handler {
	return cors.New(
		cors.Config{
			AllowOrigins:  "*",
			AllowMethods:  "GET,HEAD,PUT,POST,DELETE,PATCH",
			AllowHeaders:  "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
			ExposeHeaders: "Content-Length",
		},
	)
}
