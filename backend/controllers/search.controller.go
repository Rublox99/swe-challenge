package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var emailsURL = "http://localhost:4080/api/emails/_search"

func GetEmails(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	size := r.URL.Query().Get("size")

	if from == "" {
		from = "0"
	}
	if size == "" {
		size = "10"
	}

	query := fmt.Sprintf(`
	{
    	"search_type": "alldocuments",
		"from": %s,
		"max_results": %s
	}`, from, size)

	req, err := http.NewRequest("POST", emailsURL, strings.NewReader(query))
	if err != nil {
		http.Error(w, "Failed to create HTTP request", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth("admin", "admin")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to ZincSearch", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from ZincSearch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func SearchEmailsWithFilters(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	size := r.URL.Query().Get("size")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	text := r.URL.Query().Get("text")

	if from == "" {
		from = "0"
	}
	if size == "" {
		size = "10"
	}
	if startDate == "" {
		startDate = "2000-01-01T00:00:00.000Z"
	}
	if endDate == "" {
		endDate = "2024-12-31T23:59:59.999Z"
	}

	query := fmt.Sprintf(`
	{
		"search_type": "match",
		"query": {
			"term": "%s",
			"fields": ["Body", "Subject", "From", "To"],
			"fuzziness": 1,
			"prefix_length": 5
		},
		"filter": {
			"range": {
				"Date": {
					"gte": "%s",
					"lte": "%s"
				}
			}
		},
		"from": %s,
		"max_results": %s
	}`, text, startDate, endDate, from, size)

	req, err := http.NewRequest("POST", emailsURL, strings.NewReader(query))
	if err != nil {
		http.Error(w, "Failed to create HTTP request", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth("admin", "admin")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to ZincSearch", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from ZincSearch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func SearchEmailByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`
	{
		"query": {
			"term": {
				"_id": "%s"
			}
		}
	}`, id)

	req, err := http.NewRequest("POST", emailsURL, strings.NewReader(query))
	if err != nil {
		http.Error(w, "Failed to create HTTP request", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth("admin", "admin")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to ZincSearch", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from ZincSearch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
