package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/thomashancock/URL-Shortener/internal/database"
	"github.com/thomashancock/URL-Shortener/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	// Make Logger
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	})

	// Declare server
	mux := http.NewServeMux()

	// Create map to store URLs
	dbFile := "./sqlite-database.db"
	db := database.NewSQLDatabase(log, dbFile)

	// Create and register a router to direct HTTP requests
	router := handlers.NewRouteHandler(log, db)
	mux.HandleFunc("/", router)

	// Run server
	go http.ListenAndServe(":8080", mux)
	log.Infoln("Listening on http://localhost:8080")

	// Run until interrupt signal received
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
