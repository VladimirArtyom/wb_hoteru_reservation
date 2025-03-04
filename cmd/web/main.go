package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/config"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/handlers"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/render"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const portNumber string = ":8080"

func main() {

	var appConfig config.AppConfig = config.AppConfig{
		UseCache: false,
	}
	render.NewTemplate(&appConfig)
	tc, err := render.CreateTemplateCache()
	if err != nil  {
		log.Panic("Could not create template Cache")
	}
	
	appConfig.TemplateCache = tc

	// handlers
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	
	fmt.Println("Starting application on port", portNumber)
	http.ListenAndServe(portNumber, mux)
	
	

}
