package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go pprof()
	go DoSomething()
	wg.Wait()
}

func DoSomething() {
	counter := 0
	for {
		counter++
		fmt.Printf("counter %d\n", counter)
		time.Sleep(1 * time.Second)
		if counter > 60 {
			wg.Done()
			return
		}
	}
}

func pprof() {
	// Debug tool -> http://localhost:8080/debug/pprof/heap
	// Debug tool -> http://localhost:8080/debug/pprof/goroutine
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("PProf error %v\n", err)
	}
}
