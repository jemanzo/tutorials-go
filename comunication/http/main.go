package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"strings"
)

func GetUrl(url string) (res *http.Response) {
	req, _ := http.NewRequest("GET", url, nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotFirstResponseByte: func() {
			fmt.Printf("GotFirstResponseByte\n")
		},
		WroteHeaderField: func(key string, value []string) {
			arr := &[]string{
				strings.Join(value, "."),
				"DONE",
			}
			fmt.Printf("%s -- %s\n", key, strings.Join(*arr, " ,, "))
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	return
}
