package models

// WebhookPayload representa la estructura raíz que envía GoWa
type WebhookPayload struct {
	Event    string      `json:"event"`
	DeviceId string      `json:"device_id"`
	Payload  MessageData `json:"payload"`
}

// MessageData contiene la información del mensaje y el remitente
type MessageData struct {
	Info struct {
		ID        string `json:"id"`
		Sender    string `json:"sender"` // JID del usuario (ej: 595981... @s.whatsapp.net)
		PushName  string `json:"pushName"`
		Timestamp int64  `json:"timestamp"`
		IsFromMe  bool   `json:"isFromMe"`
	} `json:"info"`
	Message struct {
		Conversation string `json:"conversation"` // El texto del mensaje
	} `json:"message"`
}
