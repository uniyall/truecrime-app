package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"prateekkuniyal.dev/desicrimepod/data"
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
	// Start Logger
	logInstance := initializeLogger()

	// Env variables
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("No .env file was found", err)
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("Invalid DB Connection String")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB %v", err)
	}
	defer db.Close()

	// Initialize Case Repository
	caseRepo, err := data.NewCaseRepo(db, logInstance)
	if err != nil {
		logInstance.Error("Failed to Initialize Case Repository", err)
	}

	// Initialize Case Handler struct
	casesHandler := handlers.NewCaseHandler(caseRepo, logInstance)
	// API Routes
	http.HandleFunc("/api/v1/cases", casesHandler.GetCases)
	http.HandleFunc("/api/v1/cases/random/", casesHandler.GetRandomCases)

	// Server Frontend from root endpoint
	http.Handle("/", http.FileServer(http.Dir("public")))

	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server Failed", err)
	}
}
