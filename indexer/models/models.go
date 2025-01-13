package models

type Email struct {
	MessageId   string   `json:"Message-Id"`
	Date        int64    `json:"Date"`
	From        string   `json:"From"`
	To          []string `json:"To"`
	Subject     string   `json:"Subject"`
	Body        string   `json:"Body"`
	CC          []string `json:"CC"`
	MimeVersion string   `json:"Mime-Version"`
	ContentType string   `json:"Content-Type"`
	XFrom       string   `json:"X-From"`
	XTo         []string `json:"X-To"`
	XCc         []string `json:"X-cc"`
	XBcc        []string `json:"X-bcc"`
	XFolder     string   `json:"X-Folder"`
	XOrigin     string   `json:"X-Origin"`
	XFileName   string   `json:"X-Filename"`
}
