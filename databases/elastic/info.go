package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func GetInfo(es *elasticsearch.Client) {
	var r map[string]interface{}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	res.IsError()
	if err != nil {
		log.Fatalf("Error: %s", res.String())
	}

	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print client and server version numbers.
	log.Println()
	log.Println(strings.Repeat("-", 37))
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("-", 37))
	log.Println()
}
