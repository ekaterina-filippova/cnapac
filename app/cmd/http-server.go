package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func hit(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}
	fmt.Fprintf(w, "You've hit %v\n", hostname)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hit", hit)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
