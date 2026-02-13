package main

import (
	"gowa-webhook-handler/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env, usando variables de sistema")
	}

	app := fiber.New()

	// Middleware de Logger para ver las peticiones que llegan en la consola
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Webhook Handler is Running!")
	})

	// Endpoint donde GoWa enviará los datos
	app.Post("/webhook", handlers.HandleGowaWebhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor Webhook iniciado en el puerto %s", port)

	// Si app.Listen falla, log.Fatal detendrá el programa y nos dirá POR QUÉ
	log.Fatal(app.Listen(":" + port))
}
