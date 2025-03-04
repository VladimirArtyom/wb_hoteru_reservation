package handlers

import (
	"net/http"

	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/config"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/models"
	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/render"
)

var Repo *Repository;

type Repository struct {
	app *config.AppConfig
} 

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		app: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	
	stringMapData := map[string]string{
		"test": "Wei wou wehuoi",
		"jamet": "kuproy jamet",
	}

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMapData,
	})
}
