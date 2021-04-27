package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	v1 "goarch/api/v1"
	"goarch/app"
	"goarch/app/repositories/mongo"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	var port int
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.IntVar(&port, "port", 8000, "the port of the service")
	flag.Parse()

	srv := app.NewServer(mux.NewRouter())

	connection, err := mongo.NewConnection("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatalf("unable to create connection due: %s", err)
	}

	v1.Register(srv.Router, connection)

	go srv.Run(fmt.Sprintf(":%d", port))

	log.Printf("service running on port %d", port)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	go srv.Stop(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
