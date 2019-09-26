package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/clementauger/practical-golang-docker/model"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! Please hit the `/qod` API to get the quote of the day."))
}

func main() {

	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + getEnv("EXPOSE", "8080"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting producer", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		consumerURL := os.Getenv("CONSUMER_URL")
		if consumerURL == "" {
			log.Println("no consumer url...")
			return
		}
		consumerURL = fmt.Sprintf("http://%v", consumerURL)
		for {
			res, err := http.Get(consumerURL)
			if err != nil {
				log.Println(err)
				<-time.After(time.Second * 4)
				continue
			}
			log.Println("consumer response")
			io.Copy(os.Stdout, res.Body)
			res.Body.Close()
			<-time.After(time.Second * 4)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
