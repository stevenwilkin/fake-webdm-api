package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)

	http.HandleFunc("/api/v2/packages/", JSONHandler)
	http.HandleFunc("/icons/", IconHandler)

	handler := NewLoggingMiddleware(http.DefaultServeMux)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		panic("Error starting!")
	}
}
