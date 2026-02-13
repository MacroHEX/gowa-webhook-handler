package services

import (
	"fmt"
	"strings"
)

// ProcessBotLogic decide qué responder basado en el mensaje y el contexto
func ProcessBotLogic(senderName, message string) string {
	cleanMsg := strings.ToLower(strings.TrimSpace(message))

	// Lógica de Menú Principal
	if cleanMsg == "menu" || cleanMsg == "hola" || cleanMsg == "1" || cleanMsg == "2" || cleanMsg == "3" {
		return handleMenu(senderName, cleanMsg)
	}

	// Respuesta por defecto si no entiende el comando
	return fmt.Sprintf("Hola *%s*, no he reconocido ese comando.\n\nEscribe *MENU* para ver las opciones disponibles.", senderName)
}

func handleMenu(name, msg string) string {
	switch msg {
	case "1":
		return "📅 *Turnos Disponibles:*\nContamos con espacios para mañana:\n- 09:00 AM\n- 14:30 PM\n\n¿Deseas agendar alguno?"
	case "2":
		return "👨‍⚕️ *Especialistas:*\n1. Dr. Medina (Clínica)\n2. Dra. González (Pediatría)\n3. Dr. Pérez (Odontología)"
	case "3":
		return "📞 *Asesor:*\nEn un momento un agente humano se comunicará contigo. Gracias por tu paciencia."
	default:
		return fmt.Sprintf("¡Hola *%s*! 👋 Bienvenido a la demo de consulta.\n\nPor favor, selecciona una opción enviando el número:\n\n1️⃣ Consultar turnos disponibles\n2️⃣ Consultar especialistas\n3️⃣ Hablar con un Asesor", name)
	}
}
