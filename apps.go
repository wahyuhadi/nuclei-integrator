package main

import (
	"flag"
	"os"

	"github.com/wahyuhadi/nuclei-integrator/models"
	"github.com/wahyuhadi/nuclei-integrator/services"
)

var (
	epush  = flag.Bool("ep", false, "Elastic push")
	eurl   = flag.String("eurl", "http://127.0.0.1:9200", "Elastic URL")
	euser  = flag.String("euser", "", "elastic users")
	epass  = flag.String("epass", "", "elastic pass")
	eindex = flag.String("eindex", "nuclei-dast", "elastic index")
	jira   = flag.Bool("jira", false, "auto create issue on jira")
)

func parseOptions() (opts *models.Options) {
	flag.Parse()
	return &models.Options{
		Elasic:      *epush,
		ElasticUrl:  *eurl,
		ElasticUser: *euser,
		ElasticPass: *epass,
		ElasicIndex: *eindex,
		Jira:        *jira,
		JiraUser:    os.Getenv("jira_user"),
		JiraToken:   os.Getenv("jira_token"),
		JiraURI:     os.Getenv("jira_url"),
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
