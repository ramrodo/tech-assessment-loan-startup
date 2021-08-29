package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ramrodo/tech-assessment-loan-startup/config"
	"github.com/ramrodo/tech-assessment-loan-startup/router"
	log "github.com/sirupsen/logrus"
)

const (
	defaultGracefulTimeout = time.Second * 15

	serverWriteTimeout = time.Second * 10
	serverReadTimeOut  = time.Second * 10
	serverIdleTimeout  = time.Second * 10
)

var port *uint
var gracefulTimeout *time.Duration

func init() {
	config.ReadConfig()

	// Config for server
	port = flag.Uint("port", uint(config.C.Server.Port), "server port")
	gracefulTimeout = flag.Duration("graceful-timeout", defaultGracefulTimeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
}

func main() {
	router := router.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		WriteTimeout: serverWriteTimeout,
		ReadTimeout:  serverReadTimeOut,
		IdleTimeout:  serverIdleTimeout,
		Handler:      router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Infof("starting server on port %d...\n", *port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), *gracefulTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Info("shutting down")
	os.Exit(0)
}
