package handlers

import (
	"fmt"
	"gowa-webhook-handler/models"
	"gowa-webhook-handler/services"

	"github.com/gofiber/fiber/v2"
)

func HandleGowaWebhook(c *fiber.Ctx) error {
	var body models.WebhookPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Ignorar mensajes enviados por nosotros mismos para evitar bucles infinitos
	if body.Payload.Info.IsFromMe {
		return c.SendStatus(200)
	}

	fmt.Printf("Mensaje recibido de %s: %s\n", body.Payload.Info.PushName, body.Payload.Message.Conversation)

	// Lógica simple de respuesta
	if body.Payload.Message.Conversation == "Hola" {
		services.SendReply(body.Payload.Info.Sender, "¡Hola! ¿En qué puedo ayudarte hoy?")
	}

	return c.SendStatus(200)
}
