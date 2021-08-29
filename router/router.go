package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ramrodo/tech-assessment-loan-startup/config"
	"github.com/ramrodo/tech-assessment-loan-startup/handler"
)

var sfnClient *sfn.SFN

func init() {
	session, err := session.NewSession(&aws.Config{
		Region: &config.C.AWS.Region,
	})

	if err != nil {
		panic(fmt.Errorf("error creating new AWS session: %s", err))
	}

	sfnClient = sfn.New(session)
}

func NewRouter(middlewares ...func(http.Handler) http.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middlewares...)

	r.Route("/", func(r chi.Router) {
		r.Post("/credit-assignment", handler.CreditAssignment)
	})

	return r
}
