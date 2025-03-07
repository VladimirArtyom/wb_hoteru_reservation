package handlers

import (
	"net/http"

	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/config"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/models"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/render"
)

var Repo *Repository;

type Repository struct {
	App *config.AppConfig
} 

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	var ipAddr string = r.RemoteAddr
	
	// Save to the current session context
	m.App.Session.Put(r.Context(), "remote_ip", ipAddr)


	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	
	stringMapData := map[string]interface{}{
		"test": "Wei wou wehuoi",
		"jamet": "kuproy jamet",
	}

	// Read the current session context
	ipAddr :=  m.App.Session.GetString(r.Context(), "remote_ip")

	stringMapData["remote_ip"] = ipAddr

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		AutreData: stringMapData,
	})
}
