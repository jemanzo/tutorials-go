package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	log.SetOutput(os.Stdout)

	es, err := elasticsearch.NewClient(CreateConfig())
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	GetInfo(es)

	// var wg sync.WaitGroup
	// wg.Add(1)
	// var chDoc = make(chan MyDocument)
	var chDocs = make(chan MyDocuments)
	var chDone = make(chan bool)
	// CreateDocumentWriterChannel(es, chDoc, chDone)
	CreateBulkDocumentWriterChannel(es, chDocs, chDone)

	docs1 := *CreateFakeMyDocuments(10)
	docs2 := *CreateFakeMyDocuments(10)

	// go func() {
	// 	for _, v := range docs {
	// 		chDoc <- v
	// 	}
	// 	close(chDoc)
	// }()

	go func() {
		chDocs <- docs1
		chDocs <- docs2
		close(chDocs)
	}()

	<-chDone
	// wg.Done()
	// wg.Wait()
}

func Search(es *elasticsearch.Client) {
	var (
		r map[string]interface{}
		// wg sync.WaitGroup
	)
	// wg.Add(1)
	// wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Title": "two",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}

func indexDocs(es *elasticsearch.Client, wg *sync.WaitGroup) {
	// 2. Index documents concurrently
	//
	for i, title := range []string{"Test One1", "Test Two2"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			data, err := json.Marshal(struct{ Title string }{Title: title})
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test-1",
				DocumentID: strconv.Itoa(i + 1),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
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
		}(i, title)
	}
}
