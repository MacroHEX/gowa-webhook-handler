package models

// WebhookPayload es la raíz del JSON de GoWa
type WebhookPayload struct {
	Event    string      `json:"event"`
	DeviceId string      `json:"device_id"`
	Payload  MessageData `json:"payload"`
}

// MessageData contiene los detalles del evento
type MessageData struct {
	Info struct {
		ID        string `json:"id"`
		Sender    string `json:"sender"`
		PushName  string `json:"pushName"`
		Timestamp int64  `json:"timestamp"`
		IsFromMe  bool   `json:"isFromMe"`
	} `json:"info"`
	// GoWa envía el contenido en este formato
	Message struct {
		Conversation string `json:"conversation"` // Mensaje de texto simple
		// Si el mensaje tiene links o formato, viene aquí:
		ExtendedTextMessage struct {
			Text string `json:"text"`
		} `json:"extendedTextMessage"`
	} `json:"message"`
}

// GetText extrae el texto sin importar el formato del mensaje
func (m *MessageData) GetText() string {
	if m.Message.Conversation != "" {
		return m.Message.Conversation
	}
	return m.Message.ExtendedTextMessage.Text
}
