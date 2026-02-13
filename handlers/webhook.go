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

	// 1. Loggear el JSON crudo para ver qué envía GoWa exactamente
	rawBody := c.Body()

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// 2. Filtrar eventos: Solo nos interesan los mensajes de texto entrantes
	if body.Event != "message" {
		//fmt.Printf("ℹ️ Evento ignorado: %s\n", body.Event)
		return c.SendStatus(200)
	}

	// 3. Evitar bucles (Mensajes enviados por el propio bot)
	if body.Payload.IsFromMe {
		return c.SendStatus(200)
	}

	// 4. Debug de datos recibidos
	textoRecibido := strings.TrimSpace(body.Payload.Body)
	nombreRemitente := body.Payload.FromName

	// Si el mensaje está vacío, imprimimos el JSON para investigar
	if textoRecibido == "" {
		fmt.Printf("❓ Mensaje vacío recibido de [%s]. JSON: %s\n", body.Payload.From, string(rawBody))
		return c.SendStatus(200)
	}

	if nombreRemitente == "" {
		nombreRemitente = "Sin Nombre"
	}

	fmt.Printf("📩 Mensaje de [%s]: %s\n", nombreRemitente, textoRecibido)

	// 5. Responder usando el nombre del remitente
	comando := strings.ToLower(textoRecibido)

	if comando == "hola" {
		mensajeRespuesta := fmt.Sprintf("¡Hola %s! 👋 Qué bueno saludarte. ¿Todo bien?", nombreRemitente)
		services.SendReply(body.Payload.From, mensajeRespuesta)
	} else if comando == "que tal todo" {
		services.SendReply(body.Payload.From, "Por aquí todo excelente, corriendo en Railway de maravilla. 🚀")
	}

	return c.SendStatus(200)
}
