package main

// This file should only include adding handlers and calling listen and serve
// handler implementations should be seperated to their own files or put in files of related handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create the server

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../public/"))))
	mux.HandleFunc("/favicon.ico", func(http.ResponseWriter, *http.Request) {})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// listen and serve
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Server error: %v\n", err.Error())
			interrupt <- syscall.SIGTERM
		}

	}()

	log.Printf("Server started on port %s\n", server.Addr[1:])

	// wait for kill signal
	<-interrupt
	log.Println("Shutting down server")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(context); err != nil {
		log.Fatalf("Server shutdown forced: %v", err)
	}

	log.Printf("Server shut down")
}
