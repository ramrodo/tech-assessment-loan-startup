package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	log "github.com/sirupsen/logrus"

	"github.com/ramrodo/tech-assessment-loan-startup/router"
	"github.com/ramrodo/tech-assessment-loan-startup/runtime"
)

const (
	defaultGracefulTimeout = time.Second * 15
	defaultPort            = 3000
	serverWriteTimeout     = time.Second * 10
	serverReadTimeOut      = time.Second * 10
	serverIdleTimeout      = time.Second * 10
)

var port int
var gracefulTimeout *time.Duration

func main() {
	flag.IntVar(&port, "port", defaultPort, "server port")

	if runtime.IsLambdaEnvironment() {
		lambda.Start(HandleLambdaEvent)
		return
	}
	startDevelopmentServer()
}

func startDevelopmentServer() {
	flag.Parse()

	gracefulTimeout = flag.Duration("graceful-timeout", defaultGracefulTimeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	router := router.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: serverWriteTimeout,
		ReadTimeout:  serverReadTimeOut,
		IdleTimeout:  serverIdleTimeout,
		Handler:      router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Infof("starting server on port %d...\n", port)
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

func HandleLambdaEvent(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	router := router.NewRouter()
	adapter := chiadapter.New(router)

	return adapter.ProxyWithContext(ctx, req)
}
