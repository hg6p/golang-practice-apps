package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshort"
	"urlshort/database"
)

func main() {

	mux := defaultMux()
	db := database.OnInit()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	dbHandler, err := urlshort.DataBaseHandler(db, mapHandler)
	if err != nil {
		log.Fatalf("error creating database handler %v", err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
