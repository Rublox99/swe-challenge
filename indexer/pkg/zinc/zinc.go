package zinc

import (
	"bytes"
	"encoding/json"
	"enron_indexer/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	zincURL  = "http://localhost:4080/api/emails/_doc"
	username = "admin"
	password = "admin"
)

// Uploads the email data to the ZincSearch
func SendToZinc(email models.Email) {
	// Marshal email to JSON for debugging
	data, err := json.MarshalIndent(email, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing to JSON: %v\n", err)
		return
	}

	// Proceed with the original logic
	req, err := http.NewRequest("POST", zincURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending data to ZincSearch: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Debug response status and body
	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("Response Status: %v\n", resp.Status)
	log.Printf("Response Body: %v\n", string(respBody))
}
