package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	lastDocID       = 0
	lastDocIDLocker sync.Mutex
)

type MyDocuments []MyDocument

type MyDocument struct {
	DocID  int     `json:"docID"`
	Field1 string  `json:"field1"`
	Field2 int32   `json:"field2"`
	Field3 float64 `json:"field3"`
	Field4 bool    `json:"field4"`
}

func NewWithFakeData() MyDocument {
	buf, _ := generateRandomBytes(16)
	f1 := fmt.Sprintf("%x", buf)
	f4 := false
	if rand.Intn(2) == 1 {
		f4 = true
	}
	return MyDocument{
		DocID:  GetNextID(),
		Field1: f1,
		Field2: rand.Int31n(3000),
		Field3: rand.Float64(),
		Field4: f4,
	}
}

func CreateFakeMyDocuments(howMany int) *MyDocuments {
	docs := make(MyDocuments, howMany)
	for i := 0; i < howMany; i++ {
		doc := NewWithFakeData()
		docs = append(docs, doc)
	}
	return &docs
}

func generateRandomBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func GetNextID() int {
	lastDocIDLocker.Lock()
	defer lastDocIDLocker.Unlock()
	lastDocID++
	return lastDocID
}
