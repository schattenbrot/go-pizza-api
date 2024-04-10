package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/explerror"
	"github.com/schattenbrot/go-pizza-api/docs"
	"github.com/schattenbrot/go-pizza-api/internal/config"
	"github.com/schattenbrot/go-pizza-api/internal/services/app"
	"github.com/schattenbrot/go-pizza-api/internal/services/pizzas"
	"github.com/schattenbrot/go-pizza-api/pkgs/responder"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Pizza API
// @version 1.0.0
// @description This is a pizza API written in Go.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.New()
	explerror.New(explerror.Options{
		Debug: true,
		Log:   config.Logger,
		SendFunction: func(w http.ResponseWriter, status int, data *explerror.Error) error {
			err := responder.Send(w, status, data)
			if err != nil {
				config.Logger.Println("something went wrong")
				return err
			}
			return nil
		},
	})

	docs.SwaggerInfo.Host = config.Domain

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedHeaders: []string{"*"},
		Debug:          config.Cors.Debug,
	}))

	// Redirect /docs to /docs/index.html
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusFound)
	})
	r.Get("/docs/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/", app.Routes())
		r.Mount("/pizzas", pizzas.Routes())
	})

	httpProtocol := "http://"
	if config.Cors.Secure {
		httpProtocol = "https://"
	}
	config.Logger.Printf("Starting server on %s%s", httpProtocol, config.Domain)
	config.Logger.Println("Routes are available at BasePath /api/v1")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r); err != nil {
		config.Logger.Println(err.Error())
	}
}
