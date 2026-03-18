package main

import (
	"log"
	"net/http"

	"prateekkuniyal.dev/desicrimepod/handlers"
	"prateekkuniyal.dev/desicrimepod/logger"
)

func initializeLogger() *logger.Logger {
	myLogger, err := logger.NewLogger("trueCrime.log")

	if err != nil {
		log.Fatalf("Failed to initialize Logger %v", err)
	}
	return myLogger
}


func main() {
	logInstance := initializeLogger()

	// Server Frontend from root endpoint
	http.Handle("/", http.FileServer(http.Dir("public")))

	// API Routes
	http.Handle("/api/v1/cases", handlers.GetCases{})
	addr := ":8080"
	err := http.ListenAndServe(addr, nil)
	
	if err != nil {
		logInstance.Error("Server Failed", err)
	}
}
