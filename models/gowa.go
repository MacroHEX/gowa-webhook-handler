package models

// WebhookPayload representa la raíz del JSON
type WebhookPayload struct {
	Event    string      `json:"event"`
	DeviceId string      `json:"device_id"`
	Payload  MessageData `json:"payload"`
}

// MessageData mapeado según tu DEBUG JSON real
type MessageData struct {
	ID        string `json:"id"`
	Body      string `json:"body"`      // Aquí viene el texto
	From      string `json:"from"`      // JID del remitente
	FromName  string `json:"from_name"` // "Martín Medina"
	ChatId    string `json:"chat_id"`
	IsFromMe  bool   `json:"is_from_me"`
	Timestamp string `json:"timestamp"`
}
type GowaResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID string `json:"id"`
	} `json:"data"`
}
