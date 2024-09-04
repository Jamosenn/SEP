package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func indexHandler(response http.ResponseWriter, request *http.Request) {
	file, err := os.Open("../frontend/index.html")
	if err != nil {
		log.Printf("error in request %q: %v\n", request.URL, err.Error())
		http.Error(response, "server error", 500)
		return
	}
	http.ServeContent(response, request, "index.html", time.Time{}, file)
}
