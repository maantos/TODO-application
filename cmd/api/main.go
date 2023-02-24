package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

const version = "1.0.0"

func main() {

	var cfg config

	//Read flags from cmd-line in none use default
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (development|staging|production)")
	flag.Parse()

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: l,
	}

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ErrorLog:     app.logger,        // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	//Below we want to allow server to gracefully shutdown
	go func() {
		app.logger.Printf("Starting %s server on port %d...", app.config.env, app.config.port)
		if err := httpServer.ListenAndServe(); err != nil {
			httpServer.ErrorLog.Fatal(err)
		}
	}()

	//Because we run server in separate go routine we need some mechanism to stop the main function from closing
	signalChnl := make(chan os.Signal, 1)
	signal.Notify(
		signalChnl,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	//receiving from channel is blocking operation, so its blocked until signal received
	sig := <-signalChnl
	log.Println("Got signal: ", sig)

	//grecefully shutdown server after 30s
	ctx, cancelShutdown := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancelShutdown()
	httpServer.Shutdown(ctx)

}
