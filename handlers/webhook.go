package handlers

import (
	"fmt"
	"gowa-webhook-handler/models"
	"gowa-webhook-handler/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleGowaWebhook(c *fiber.Ctx) error {
	fmt.Println("DEBUG JSON:", string(c.Body()))

	var body models.WebhookPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// 1. Evitar bucles: Ignorar si el mensaje lo enviamos nosotros
	if body.Payload.Info.IsFromMe {
		return c.SendStatus(200)
	}

	// 2. Extraer el texto limpio
	textoRecibido := body.Payload.GetText()
	remitente := body.Payload.Info.PushName
	if remitente == "" {
		remitente = body.Payload.Info.Sender
	}

	fmt.Printf("📩 Mensaje de [%s]: %s\n", remitente, textoRecibido)

	// 3. Lógica de respuesta inteligente
	comando := strings.ToLower(strings.TrimSpace(textoRecibido))

	switch comando {
	case "hola", "buenos días":
		services.SendReply(body.Payload.Info.Sender, "¡Hola Martín! Soy tu asistente en Go. ¿Cómo va el vlog hoy?")
	case "status":
		services.SendReply(body.Payload.Info.Sender, "✅ El sistema está operando normalmente en Railway.")
	}

	return c.SendStatus(200)
}
