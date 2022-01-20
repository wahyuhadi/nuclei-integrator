package models

import "time"

type Nuclei struct {
	TemplateID string `json:"template-id"`
	Info       struct {
		Name        string      `json:"name"`
		Author      []string    `json:"author"`
		Tags        []string    `json:"tags"`
		Description string      `json:"description"`
		Reference   interface{} `json:"reference"`
		Severity    string      `json:"severity"`
	} `json:"info"`
	Type          string    `json:"type"`
	Host          string    `json:"host"`
	MatchedAt     string    `json:"matched-at"`
	IP            string    `json:"ip"`
	Timestamp     time.Time `json:"timestamp"`
	CurlCommand   string    `json:"curl-command"`
	MatcherStatus bool      `json:"matcher-status"`
}
