package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/maantos/todoApplication/pkg/application"
	"github.com/maantos/todoApplication/pkg/db"
	thttp "github.com/maantos/todoApplication/pkg/http"
)

func main() {

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	r := chi.NewRouter()
	db := db.NewTasksStorage()
	service := application.NewTaskService(db)
	h := thttp.NewHandler(service)
	thttp.Routes(r, h)

	httpServer := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	//Below we want to allow server to gracefully shutdown
	go func() {
		l.Println("Starting server on port :8080...")
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
