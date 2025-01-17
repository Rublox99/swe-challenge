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
	username = "admin"
	password = "admin"
	ec2      = "http://ec2-52-14-213-114.us-east-2.compute.amazonaws.com:4080"
)

var zincURL = fmt.Sprintf("%s/api/emails/_doc", ec2)

func SendToZinc(email models.Email) {
	data, err := json.MarshalIndent(email, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing to JSON: %v\n", err)
		return
	}

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

	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("Response Status: %v\n", resp.Status)
	log.Printf("Response Body: %v\n", string(respBody))
}
