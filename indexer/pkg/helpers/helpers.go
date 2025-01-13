package helpers

import (
	"bufio"
	"enron_indexer/models"
	"enron_indexer/pkg/zinc"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// reads and processes a single email .txt
func ProcessFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading the file %s: %v\n", path, err)
		return
	}

	email := parseEmail(content)
	zinc.SendToZinc(email)
}

// parses the content of the email and creates an Email object
func parseEmail(content []byte) models.Email {
	email := models.Email{
		MessageId:   "*",
		Date:        time.Now().UnixNano() / int64(time.Millisecond), // Default to current time in epochMillis
		From:        "*",
		To:          []string{"*"},
		Subject:     "*",
		Body:        "",
		CC:          []string{"*"},
		MimeVersion: "*",
		ContentType: "*",
		XFrom:       "*",
		XTo:         []string{},
		XCc:         []string{},
		XBcc:        []string{},
		XFolder:     "*",
		XOrigin:     "*",
		XFileName:   "*",
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	var bodyStart bool

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "X-FileName:") {
			bodyStart = true
			continue
		}

		if bodyStart {
			email.Body += line + "\n"
			continue
		}

		switch {
		case strings.HasPrefix(line, "Message-ID:"):
			email.MessageId = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID: "))
		case strings.HasPrefix(line, "Date:"):
			dateStr := strings.TrimSpace(strings.TrimPrefix(line, "Date: "))

			// Try parsing the date in RFC1123 format
			parsedDate, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700 (MST)", dateStr)
			if err != nil {
				// If parsing fails, log the error and use the current date
				fmt.Printf("Error parsing date: %v, using current date.\n", err)
				email.Date = time.Now().UnixNano() / int64(time.Millisecond) // Default to current time
			} else {
				// Successfully parsed RFC1123 date, convert to epoch milliseconds
				email.Date = parsedDate.UnixNano() / int64(time.Millisecond)
			}
		case strings.HasPrefix(line, "From:"):
			email.From = strings.TrimSpace(strings.TrimPrefix(line, "From: "))
		case strings.HasPrefix(line, "To:"):
			email.To = parseCommaSeparated(line, "To: ")
		case strings.HasPrefix(line, "Subject:"):
			email.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject: "))
		case strings.HasPrefix(line, "Mime-Version:"):
			email.MimeVersion = strings.TrimSpace(strings.TrimPrefix(line, "Mime-Version: "))
		case strings.HasPrefix(line, "Content-Type:"):
			email.ContentType = strings.TrimSpace(strings.TrimPrefix(line, "Content-Type: "))
		case strings.HasPrefix(line, "X-From:"):
			email.XFrom = strings.TrimSpace(strings.TrimPrefix(line, "X-From: "))
		case strings.HasPrefix(line, "X-To:"):
			email.XTo = parseCommaSeparated(line, "X-To: ")
		case strings.HasPrefix(line, "X-cc:"):
			email.XCc = parseCommaSeparated(line, "X-cc: ")
		case strings.HasPrefix(line, "X-bcc:"):
			email.XBcc = parseCommaSeparated(line, "X-bcc: ")
		case strings.HasPrefix(line, "X-Folder:"):
			email.XFolder = strings.TrimSpace(strings.TrimPrefix(line, "X-Folder: "))
		case strings.HasPrefix(line, "X-Origin:"):
			email.XOrigin = strings.TrimSpace(strings.TrimPrefix(line, "X-Origin: "))
		case strings.HasPrefix(line, "X-FileName:"):
			email.XFileName = strings.TrimSpace(strings.TrimPrefix(line, "X-FileName: "))
		}
	}

	if email.Body == "" {
		email.Body = "No body content"
	}

	return email
}

// returns a strings array based on a comma-separated string
func parseCommaSeparated(line, prefix string) []string {
	field := strings.TrimPrefix(line, prefix)
	values := strings.Split(field, ",")

	for i, value := range values {
		values[i] = strings.TrimSpace(value)
	}

	return values
}
