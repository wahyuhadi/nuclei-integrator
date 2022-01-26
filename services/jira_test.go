package services

import (
	"os"
	"testing"

	"github.com/wahyuhadi/nuclei-integrator/models"
)

func Test_jiraAuth(t *testing.T) {
	t.Run("testing with valid credential", func(t *testing.T) {
		opts := &models.Options{
			JiraUser:  os.Getenv("jira_user"),
			JiraToken: os.Getenv("jira_token"),
			JiraURI:   os.Getenv("jira_url"),
		}
		gotClient, err := jiraAuth(opts)
		if err != nil {
			t.Errorf("Expected success login")
			t.Fail()
		}

		_, _, err = gotClient.User.GetSelf()
		if err != nil {
			t.Errorf("Expected get profile")
			t.Fail()
		}
	})

	t.Run("testing with invalid credential", func(t *testing.T) {
		opts := &models.Options{
			JiraUser:  "this is no username",
			JiraToken: "thisis no password",
			JiraURI:   os.Getenv("jira_url"),
		}
		gotClient, err := jiraAuth(opts)
		_, _, err = gotClient.User.GetSelf()
		if err == nil {
			t.Errorf("Expected invalid get profile")
			t.Fail()
		}
	})

	t.Run("testing with invalid Url", func(t *testing.T) {
		opts := &models.Options{
			JiraUser:  os.Getenv("jira_user"),
			JiraToken: os.Getenv("jira_token"),
			JiraURI:   "https://invalidurl.com",
		}
		gotClient, err := jiraAuth(opts)
		_, _, err = gotClient.User.GetSelf()
		if err == nil {
			t.Errorf("Expected invalid get profile")
			t.Fail()
		}
	})
}
