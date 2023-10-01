package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webstroy/main/utility"
)

type Routes struct {
	Title   string
	Handler func(http.ResponseWriter, *http.Request)
}

func CreateServer(routes []Routes) {
	mux := http.NewServeMux()

	a := utility.Intro{Title: "intro", Story: []string{}, Options: []utility.Option{}}
	mux.HandleFunc("/", utility.CreateHandler(utility.CreateTemplate(utility.IndexPage), a))
	for _, v := range routes {

		mux.HandleFunc("/"+v.Title, v.Handler)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", Logger(mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and path
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		// Measure the request processing time
		startTime := time.Now()
		defer func() {
			elapsedTime := time.Since(startTime)
			log.Printf("Completed in %s", elapsedTime)
		}()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
