package handlers

import (
	"fmt"
	"gowa-webhook-handler/models"
	"gowa-webhook-handler/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleGowaWebhook(c *fiber.Ctx) error {
	var body models.WebhookPayload

	// Debug opcional para ver el JSON en consola
	// fmt.Println("DEBUG JSON:", string(c.Body()))

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// 1. Evitar bucles infinitos
	if body.Payload.IsFromMe {
		return c.SendStatus(200)
	}

	// 2. Extraer datos reales
	textoRecibido := body.Payload.Body
	nombreRemitente := body.Payload.FromName
	if nombreRemitente == "" {
		nombreRemitente = "Desconocido"
	}

	fmt.Printf("📩 Mensaje de [%s]: %s\n", nombreRemitente, textoRecibido)

	// 3. Lógica de respuesta
	comando := strings.ToLower(strings.TrimSpace(textoRecibido))

	if comando == "hola" {
		services.SendReply(body.Payload.From, "¡Hola Martín! Recibí tu mensaje: "+textoRecibido)
	}

	return c.SendStatus(200)
}
