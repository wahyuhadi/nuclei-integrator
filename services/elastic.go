package services

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"errors"
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

func Elastic(opts *models.Options) error {
	client, err := elasticAuth(opts)
	if err != nil {
		gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error auth")
		return err
	}
	_, err = client.Info()
	if err != nil {
		gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error auth")
		return err
	}
	pushdata(opts, client)
	return nil

}

func elasticAuth(opts *models.Options) (*elasticsearch.Client, error) {

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
	c, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.New("error auths")
	}
	return c, nil
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
			continue
		}
		gologger.Info().Str("Is Success ", fmt.Sprintf("%v", response.IsError())).Msg("Success Push data to elastic")
	}
}
