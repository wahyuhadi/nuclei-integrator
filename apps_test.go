package main

import (
	"os"
	"testing"
)

func Test_parseOptions(t *testing.T) {
	t.Run("Check .env for Jira Token", func(t *testing.T) {
		token := os.Getenv("jira_token")
		if token == "" {
			t.Errorf("Expected please add jira_token to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check .env for Jira URI", func(t *testing.T) {
		token := os.Getenv("jira_url")
		if token == "" {
			t.Errorf("Expected please add jira_url to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check .env for Jira User", func(t *testing.T) {
		token := os.Getenv("jira_user")
		if token == "" {
			t.Errorf("Expected please add jira_user to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check .env for Elastic URI", func(t *testing.T) {
		token := os.Getenv("elastic_url")
		if token == "" {
			t.Errorf("Expected please add elastic_url to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check .env for Elastic user", func(t *testing.T) {
		token := os.Getenv("elastic_user")
		if token == "" {
			t.Errorf("Expected please add elastic_user to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check .env for Elastic password", func(t *testing.T) {
		token := os.Getenv("elastic_password")
		if token == "" {
			t.Errorf("Expected please add elastic_password to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check .env for Elastic Index", func(t *testing.T) {
		token := os.Getenv("elastic_index")
		if token == "" {
			t.Errorf("Expected please add elastic_index to .bashrc")
			t.Fail()
		}
	})

	t.Run("Check Jira Token", func(t *testing.T) {
		opts := parseOptions()
		if opts.JiraToken != os.Getenv("jira_token") {
			t.Errorf("Expected jira token same value with .env")
			t.Fail()
		}
	})

	t.Run("Check Jira URI", func(t *testing.T) {
		opts := parseOptions()
		if opts.JiraURI != os.Getenv("jira_url") {
			t.Errorf("Expected jira token same value with .env")
			t.Fail()
		}
	})

	t.Run("Check Jira User", func(t *testing.T) {
		opts := parseOptions()
		if opts.JiraUser != os.Getenv("jira_user") {
			t.Errorf("Expected jira token same value with .env")
			t.Fail()
		}
	})

	t.Run("Check Elastic URL", func(t *testing.T) {
		opts := parseOptions()
		if opts.ElasticUrl != os.Getenv("elastic_url") {
			t.Errorf("Expected Elastic url same value with .env")
			t.Fail()
		}
	})

	t.Run("Check Elastic user", func(t *testing.T) {
		opts := parseOptions()
		if opts.ElasticUser != os.Getenv("elastic_user") {
			t.Errorf("Expected Elastic user same value with .env")
			t.Fail()
		}
	})

	t.Run("Check Elastic password", func(t *testing.T) {
		opts := parseOptions()
		if opts.ElasticPass != os.Getenv("elastic_password") {
			t.Errorf("Expected Elastic password same value with .env")
			t.Fail()
		}
	})

	t.Run("Check Elastic index", func(t *testing.T) {
		opts := parseOptions()
		if opts.ElasicIndex != os.Getenv("elastic_index") {
			t.Errorf("Expected Elastic indexs same value with .env")
			t.Fail()
		}
	})
}
