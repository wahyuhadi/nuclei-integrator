package main

import (
	"flag"

	"github.com/wahyuhadi/nuclei-integrator/models"
	"github.com/wahyuhadi/nuclei-integrator/services"
)

var (
	epush  = flag.Bool("ep", true, "Elastic push")
	eurl   = flag.String("eurl", "http://127.0.0.1:9200", "Elastic URL")
	euser  = flag.String("euser", "", "elastic users")
	epass  = flag.String("epass", "", "elastic pass")
	eindex = flag.String("eindex", "nuclei-dast", "elastic index")
)

func parseOptions() (opts *models.Options) {
	flag.Parse()
	return &models.Options{
		Elasic:      *epush,
		ElasticUrl:  *eurl,
		ElasticUser: *euser,
		ElasticPass: *epass,
		ElasicIndex: *eindex,
	}
}

func main() {
	opts := parseOptions()
	if opts.Elasic {
		services.Elastic(opts)
	}
}
