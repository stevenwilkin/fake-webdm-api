package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	stylesheets, err := filepath.Glob(filepath.Join("www", "public", "css", "*.css"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, s := range stylesheets {
		stylesheets[i] = strings.TrimPrefix(s, "www")
	}

	htmlPath := filepath.Join("www", "templates", "index.html")
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Stylesheets []string
	}{
		Stylesheets: stylesheets,
	}
	t.Execute(w, data)
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)

	http.HandleFunc("/api/v2/packages/", JSONHandler)
	http.HandleFunc("/icons/", IconHandler)
	http.Handle("/public/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/", mainPageHandler)

	handler := NewLoggingMiddleware(http.DefaultServeMux)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		panic("Error starting!")
	}
}
