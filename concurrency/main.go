package main

import (
	"fmt"
	"sync"
)

var (
	counter     = 0
	counterLock = sync.Mutex{}
	wg          sync.WaitGroup
)

func main() {
	runCounter()
}

func runCounter() {
	total := 10000
	fmt.Printf("Starting %d goroutines!\n", total)
	// wg.Add(total)
	for i := 0; i < total; i++ {
		wg.Add(1)
		go counterRoutine()
		// go counterInc()
	}
	fmt.Printf("Counter(%d) before sleep!\n", counter)
	// time.Sleep(5 * time.Second)
	wg.Wait()
	fmt.Printf("Counter(%d) after sleep!\n", counter)
}

func counterRoutine() {
	counterInc()
}

func counterInc() {
	defer wg.Done()
	// counterLock.Lock()
	counter++
	// counterLock.Unlock()
}

func runJobs() {
	jobs := make(chan string, 4)
	worker1(jobs)
	go worker2(jobs)
	worker1(jobs)
	for i := 0; i < 6; i++ {
		fmt.Println(<-jobs)
	}
}

func worker1(c chan string) {
	c <- "worker1"
	c <- "worker1"
}

func worker2(c chan string) {
	c <- "worker2"
	c <- "worker2"
}
