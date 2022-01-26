package models

import "time"

type Nuclei struct {
	Template    string `json:"template"`
	TemplateURL string `json:"template-url"`
	TemplateID  string `json:"template-id"`
	Info        struct {
		Name           string   `json:"name"`
		Author         []string `json:"author"`
		Tags           []string `json:"tags"`
		Description    string   `json:"description"`
		Reference      []string `json:"reference"`
		Severity       string   `json:"severity"`
		Classification struct {
			CveID       []string `json:"cve-id"`
			CweID       []string `json:"cwe-id"`
			CvssMetrics string   `json:"cvss-metrics"`
			CvssScore   float64  `json:"cvss-score"`
		} `json:"classification"`
	} `json:"info"`
	Type             string   `json:"type"`
	Host             string   `json:"host"`
	MatchedAt        string   `json:"matched-at"`
	CurlCommand      string   `json:"curl-command"`
	ExtractedResults []string `json:"extracted-results"`
	Meta             struct {
		Hostname string `json:"Hostname"`
	} `json:"meta"`
	IP            string    `json:"ip"`
	Timestamp     time.Time `json:"timestamp"`
	MatcherStatus bool      `json:"matcher-status"`
}
