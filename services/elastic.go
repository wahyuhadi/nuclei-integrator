package services

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/ESgo/es"
	"github.com/wahyuhadi/nuclei-integrator/models"
)

// Model data

func Elastic(opts *models.Options) {
	cfg := elasticsearch.Config{
		Addresses: []string{opts.ElasticUrl},
		Username:  opts.ElasticUser, // if ES need this
		Password:  opts.ElasticPass,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
		},
	}
	c, _ := elasticsearch.NewClient(cfg)
	// PushexamplePushData(c)
	pushdata(opts, c)

}

func pushdata(opts *models.Options, c *elasticsearch.Client) {
	sc := bufio.NewScanner(os.Stdin)
	var nuclei models.Nuclei
	for sc.Scan() {
		// err := json.NewDecoder(os.Stdin).Decode(&nuclei)
		// fmt.Println(sc.Text())
		json.Unmarshal([]byte(sc.Text()), &nuclei)

		// parsing with esutil from elastic
		data := esutil.NewJSONReader(&nuclei)
		// Push data to elastic
		response, err := es.PushData(c, opts.ElasicIndex, data)
		if err != nil {
			gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error push data")
		}
		gologger.Info().Str("Is Success ", fmt.Sprintf("%v", response.IsError())).Msg("Success Push data to elastic")
	}
}
