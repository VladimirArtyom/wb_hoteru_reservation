package main

import (
	"net/http"

	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/config"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/handlers"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewMux()

	//mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad) // Must preseve this one to be able to interact with the data in this current session

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
