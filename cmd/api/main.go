package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handlers "github.com/cholaraja/revise/goweb/basicStructure/cmd/api/handlers/products"
	"github.com/cholaraja/revise/goweb/basicStructure/internal/platform/database"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatalf("error: Connecting to Database: %s", err)
	}
	defer db.Close()

	ProductHandlers := handlers.Product{DB: db}

	server := http.Server{
		Addr:    ":8000",
		Handler: http.HandlerFunc(ProductHandlers.ListProducts),
	}

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()
	log.Printf("Server started at Port %s, Press Ctrl+C to cancel", server.Addr)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("error: Listening and Serving: %s", err)
	case <-osSignals:
		log.Print("Caught Signal, Shutting down")

		const timeout = time.Second * 15
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("error: Gracefully shutting down server: %s", err)
			if err := server.Close(); err != nil {
				log.Fatalf("error: closing server: %s", err)
			}
		}
	}
}
