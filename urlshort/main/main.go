package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"urlshort"
	myTypes "urlshort/types"
)

func main() {
	mux := defaultMux()
	jsonFilePath := flag.String("json", "", "Path to the JSON file")
	flagFilePath := flag.String("yaml", "../urlshort.yaml", "Path to the YAML file")
	flag.Parse()
	var data map[string]myTypes.T
	if *jsonFilePath != "" {

		data = urlshort.ParseJson(*jsonFilePath)
		// Use the YAML configuration as needed
		fmt.Printf("JSON Config: %+v\n", data)
	} else {
		data = urlshort.CreateMap(urlshort.ParseYaml(*flagFilePath))

		// Use the JSON configuration as needed
		fmt.Printf("YAML Config: %+v\n", data)
	}
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(data, mapHandler)

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
