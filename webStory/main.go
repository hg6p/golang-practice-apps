package main

import (
	"fmt"
	server "webstroy/main/server"
	"webstroy/main/utility"
)

func main() {

	formmatedJson := utility.ParseJson()

	htmlTemplate := utility.CreateTemplate(utility.StoryPages)
	if htmlTemplate == nil {
		fmt.Println("Html template in null")
		return
	}

	var routes []server.Routes
	for val, v := range formmatedJson {
		handler := utility.CreateHandler(htmlTemplate, v)
		routes = append(routes, server.Routes{Title: val, Handler: handler})
	}
	server.CreateServer(routes)

}
