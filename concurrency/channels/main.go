package main

import (
	"log"
	"os"
	"sync"
)

const CHANNELS_PRINT = "\nChannels\n  A len(%d) cap(%d)\n  B len(%d) cap(%d)\n\n"

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime)

	RunNilChannel()
	// RunChannels()

	log.Println("All done, thanks!")
}

func RunNilChannel() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch1 <- "Hi1"
	ch2 <- "Hi2"
	close(ch1)
	close(ch2)
	// for i := 0; i < 4; i++ {
	for ch1 != nil || ch2 != nil {
		select {
		case val, ok := <-ch1:
			log.Println(val, ok)
			ch1 = nil
		case val, ok := <-ch2:
			log.Println(val, ok)
			ch2 = nil
		}
	}
}

func RunChannels() {
	var wg sync.WaitGroup

	listA := []string{"A1", "A2", "A3", "A4", "A5", "A6"}
	listB := []string{"B1", "B2", "B3", "B4", "B5", "B6"}

	totalWriters := 2
	totalReaders := 40

	// Unbuffered channels
	msgA := make(chan string)
	msgB := make(chan string)

	// Buffered channels
	// msgA := make(chan string, 2)
	// msgB := make(chan string, 2)

	// log.Printf(CHANNELS_PRINT, len(msgA), cap(msgA), len(msgB), cap(msgB))

	log.Printf("Creating %d Writers\n", totalWriters)
	wg.Add(2)
	go Write(&wg, listA, msgA)
	go Write(&wg, listB, msgB)

	log.Printf("Creating %d Readers\n", totalReaders)
	for i := 0; i < totalReaders; i++ {
		wg.Add(1)
		go Read(&wg, i, msgA, msgB)
	}

	log.Println("Waiting ...")
	wg.Wait()
}

func Read(wg *sync.WaitGroup, id int, msgA, msgB <-chan string) {
	defer wg.Done()

	log.Printf("[%d] Reader goroutine started\n", id)
	defer log.Printf("[%d] Reader goroutine finished\n", id)

	var (
		loopCounter int  = 0
		okA, okB    bool = true, true
		gotA, gotB  string
	)

	// for {
	// for okA || okA {
	for msgA != nil || msgB != nil {
		loopCounter++

		// hold := time.Duration(rand.Intn(6))
		// log.Printf("[%d] select(%d) will read in %d sec(s)\n", id, loopCounter, hold)
		// time.Sleep(hold * time.Second)

		// log.Printf(CHANNELS_PRINT, len(msgA), cap(msgA), len(msgB), cap(msgB))
		select {
		case gotA, okA = <-msgA:
			log.Printf("[%d] select(%d) Channel A is open? %v\n", id, loopCounter, okA)
			if !okA {
				msgA = nil
				continue
			}
			log.Printf("[%d] select(%d) Channel A got %q\n", id, loopCounter, gotA)
		case gotB, okB = <-msgB:
			log.Printf("[%d] select(%d) Channel B is open? %v\n", id, loopCounter, okB)
			if !okB {
				msgB = nil
				continue
			}
			log.Printf("[%d] select(%d) Channel B got %q\n", id, loopCounter, gotB)
		}
	}
}

func Write(wg *sync.WaitGroup, arr []string, msg chan<- string) {
	id := string(arr[0][0])

	log.Printf("[%s] Writer goroutine started\n", id)

	defer func() {
		wg.Done()
		log.Printf("[%s] Writer goroutine finished\n", id)
		log.Printf("[%s] Writer closing channel %s\n", id, id)
		close(msg)
	}()

	for _, item := range arr {
		// hold := time.Duration(rand.Intn(5))
		// log.Printf("[%s]   will write %q in %d sec(s)\n", id, item, hold)
		// time.Sleep(hold * time.Second)

		msg <- item
	}
}
