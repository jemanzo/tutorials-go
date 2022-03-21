package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func CreateDocumentWriterChannel(es *elasticsearch.Client, ChReader <-chan MyDocument, ChDone chan<- bool) {
	log.Println("Create Document Writer Channel")
	go func() {
		for {
			myDoc, ok := <-ChReader
			if !ok {
				ChDone <- true
				return
			}
			log.Println(myDoc)

			// Build the request body.
			data, err := json.Marshal(myDoc)
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "tutorial-1",
				DocumentID: strconv.FormatInt(int64(myDoc.DocID), 10),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), myDoc.DocID)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
			res.Body.Close()
		}
	}()
}

func CreateBulkDocumentWriterChannel(es *elasticsearch.Client, ChReader <-chan MyDocuments, ChDone chan<- bool) {
	log.Println("Create Document Writer Channel")
	go func() {
		for {
			myDoc, ok := <-ChReader
			if !ok {
				ChDone <- true
				return
			}
			log.Println(myDoc)

			// Build the request body.
			data, err := json.Marshal(myDoc)
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.BulkRequest{
				Index:   "tutorial-2",
				Body:    bytes.NewReader(data),
				Refresh: "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}

			if res.IsError() {
				log.Printf("[%s] Error %v", res.Status(), err)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
			res.Body.Close()
		}
	}()
}
