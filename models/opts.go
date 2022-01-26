package models

type Options struct {
	Elasic      bool
	ElasticUrl  string
	ElasticUser string
	ElasticPass string
	ElasicIndex string

	Jira      bool
	JiraUser  string
	JiraToken string
	JiraURI   string
}
