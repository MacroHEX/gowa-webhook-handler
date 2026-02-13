package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendReply(to string, text string) {
	baseURL := os.Getenv("GOWA_API_URL")
	user := os.Getenv("GOWA_BASIC_AUTH_USER")
	pass := os.Getenv("GOWA_BASIC_AUTH_PASS")

	apiURL := fmt.Sprintf("%s/send/message", baseURL)

	payload := map[string]interface{}{
		"phone":   to,
		"message": text,
	}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error creando request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error contactando a GoWa: %v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Printf("Respuesta enviada a %s, Status: %s\n", to, resp.Status)
}
