package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/config"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/handlers"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"
var appConfig config.AppConfig = config.AppConfig{
	UseCache: false,
}
var session *scs.SessionManager


func setSession(ses *scs.SessionManager) *scs.SessionManager {
	ses.Lifetime = 24 * time.Hour
	ses.Cookie.Secure = false
	ses.Cookie.SameSite = http.SameSiteLaxMode
	ses.Cookie.Persist = true
	return ses
}

func main() {

	session = scs.New()
	session = setSession(session)
	// Put session into the appConfig to make it available in another package

	render.NewTemplate(&appConfig)
	tc, err := render.CreateTemplateCache()
	if err != nil  {
		log.Panic("Could not create template Cache", err)
	}
	
	appConfig.TemplateCache = tc
	appConfig.Session = session

	// set the handlers repostiory
	repo := handlers.NewRepository(&appConfig)
	handlers.NewHandlers(repo)
	render.NewTemplate(&appConfig)

	// handlers
	fmt.Println("Starting application on port", portNumber)
	http.ListenAndServe(portNumber, routes(&appConfig))

}
