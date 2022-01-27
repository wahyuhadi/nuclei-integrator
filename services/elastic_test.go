package services

import (
	"os"
	"testing"

	"github.com/wahyuhadi/nuclei-integrator/models"
)

func Test_authElastic(t *testing.T) {
	t.Run("testing with invalid elastic host", func(t *testing.T) {
		opts := &models.Options{
			ElasticUrl: "https://0.0.0.1",
		}
		c, _ := elasticAuth(opts)
		_, err := c.Info()
		if err == nil {
			t.Errorf("Expected no route to host")
			t.Fail()
		}
	})

	t.Run("testing with valid elastic env", func(t *testing.T) {
		opts := &models.Options{
			ElasticUrl:  os.Getenv("elastic_url"),
			ElasicIndex: os.Getenv("elastic_index"),
			ElasticUser: os.Getenv("elastic_user"),
			ElasticPass: os.Getenv("elastic_password"),
		}
		c, _ := elasticAuth(opts)
		_, err := c.Info()
		if err != nil {
			t.Errorf("Invalid Elastic environtment ->> Expected get sessions")
			t.Errorf("Please check your .bashrc")
			t.Fail()
		}
	})
}

func Test_Elastic(t *testing.T) {
	t.Run("testing with invalid elastic host", func(t *testing.T) {
		opts := &models.Options{
			ElasticUrl: "https://0.0.0.1",
		}
		err := Elastic(opts)
		if err == nil {
			t.Errorf("Expected no route to host")
			t.Fail()
		}
	})

	t.Run("testing with valid elastic env", func(t *testing.T) {
		opts := &models.Options{
			ElasticUrl:  os.Getenv("elastic_url"),
			ElasicIndex: os.Getenv("elastic_index"),
			ElasticUser: os.Getenv("elastic_user"),
			ElasticPass: os.Getenv("elastic_password"),
		}
		err := Elastic(opts)
		if err != nil {
			t.Errorf("Invalid Elastic environtment ->> Expected get sessions")
			t.Errorf("Please check your .bashrc")
			t.Fail()
		}
	})
}
