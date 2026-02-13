package handlers

import (
	"gowa-webhook-handler/models"
	"gowa-webhook-handler/services"

	"github.com/gofiber/fiber/v2"
)

func HandleGowaWebhook(c *fiber.Ctx) error {
	var body models.WebhookPayload
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	// Filtros de seguridad
	if body.Event != "message" || body.Payload.IsFromMe || body.Payload.Body == "" {
		return c.SendStatus(200)
	}

	nombre := body.Payload.FromName
	if nombre == "" {
		nombre = "Cliente"
	}

	// Obtener la respuesta de nuestra lógica de negocio
	respuestaTexto := services.ProcessBotLogic(nombre, body.Payload.Body)

	// Enviar de vuelta a WhatsApp
	if respuestaTexto != "" {
		go services.SendReply(body.Payload.From, respuestaTexto)
	}

	return c.SendStatus(200)
}
