package main

import (
	"net/http"

	"github.com/BuzzBumble/alwaysallin/config"
	"github.com/BuzzBumble/alwaysallin/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(GenerateCSRF)
	mux.Use(middleware.Recoverer)

	mux.Use(SessionLoad)

	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
