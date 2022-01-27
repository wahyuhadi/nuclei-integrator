package services

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	jira "github.com/andygrunwald/go-jira"
	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/nuclei-integrator/models"
)

func Jira(opts *models.Options) {
	client, err := jiraAuth(opts)
	if err != nil {
		gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error create issue")
		return
	}

	createIssue(client)

}

func createIssue(client *jira.Client) {
	sc := bufio.NewScanner(os.Stdin)
	var nuclei models.Nuclei
	for sc.Scan() {
		json.Unmarshal([]byte(sc.Text()), &nuclei)
		i := jira.Issue{
			Fields: &jira.IssueFields{
				Description: createSummary(&nuclei),
				Type: jira.IssueType{
					Name: "Task",
				},
				Project: jira.Project{
					Key: "ST",
				},
				Summary: nuclei.Info.Name,
			},
		}
		issue, _, err := client.Issue.Create(&i)
		if err != nil {
			gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error create issue")
			continue
		}
		gologger.Info().Str("Is Success ", fmt.Sprintf("%s", issue.Key)).Msg("Success create issue")

	}

}

func createSummary(nuclei *models.Nuclei) string {
	Desc := fmt.Sprintf("Desc : %s\n", nuclei.Info.Description)
	Ip := fmt.Sprintf("IP : %s\n", nuclei.IP)
	Host := fmt.Sprintf("Host : %v\n", nuclei.Host)
	Severity := fmt.Sprintf("Severity : %s\n", nuclei.Info.Severity)
	Ref := fmt.Sprintf("Ref : %s\n", nuclei.Info.Reference)
	Res := fmt.Sprintf("Result : %v\n", nuclei.ExtractedResults)
	Curl := fmt.Sprintf("Curl : %v\n", nuclei.CurlCommand)

	summ := Desc + Ip + Host + Severity + Ref + Res + Curl
	return summ
}

func jiraAuth(opts *models.Options) (client *jira.Client, err error) {
	// Create a BasicAuth Transport object
	tp := jira.BasicAuthTransport{
		Username: opts.JiraUser,
		Password: opts.JiraToken,
	}
	// Create a new Jira Client
	client, err = jira.NewClient(tp.Client(), opts.JiraURI)
	if err != nil {
		return client, errors.New("Invalid credentials")
	}
	return client, nil
}
