package main

import (
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func CreateConfig() elasticsearch.Config {
	cacert := "./certs/es01-chain.pem"
	cert, err := ioutil.ReadFile(cacert)
	if err != nil {
		log.Fatalf("Elasticsearch Certificate: ReadError file %s \n%v", cacert, err)
	}

	return elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "<password>",
		CACert:   cert,
	}
}
