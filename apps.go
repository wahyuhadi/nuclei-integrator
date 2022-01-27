package main

import (
	"flag"
	"os"

	"github.com/wahyuhadi/nuclei-integrator/models"
	"github.com/wahyuhadi/nuclei-integrator/services"
)

var (
	epush = flag.Bool("ep", false, "Elastic push")
	jira  = flag.Bool("jira", false, "auto create issue on jira")
)

func parseOptions() (opts *models.Options) {
	flag.Parse()
	return &models.Options{
		Elasic:      *epush,
		ElasticUrl:  os.Getenv("elastic_url"),
		ElasticUser: os.Getenv("elastic_user"),
		ElasticPass: os.Getenv("elastic_password"),
		ElasicIndex: os.Getenv("elastic_index"),

		// Jira part
		Jira:      *jira,
		JiraUser:  os.Getenv("jira_user"),
		JiraToken: os.Getenv("jira_token"),
		JiraURI:   os.Getenv("jira_url"),
	}
}

func main() {
	opts := parseOptions()

	// Push data to elastic
	if opts.Elasic {
		services.Elastic(opts)
	}

	// auto create issue in Jira
	if opts.Jira {
		services.Jira(opts)
	}
}
