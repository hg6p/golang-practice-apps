package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"urlshort"
)

func main() {
	mux := defaultMux()
	flagFilePath := flag.String("yaml", "urlshort.yaml", "Path to the YAML file")
	flag.Parse()
	println(*flagFilePath)

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yamlData := urlshort.CreateMap(urlshort.ParseYaml(*flagFilePath))
	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)

	if err != nil {
		log.Fatalf("error creating YAML handler: %v", err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
