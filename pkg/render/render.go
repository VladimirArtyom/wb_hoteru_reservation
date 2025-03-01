package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/VladimirArtyom/wb_hoteru_reservation/pkg/config"
)

var app *config.AppConfig

func NewTemplate(app_param *config.AppConfig) {
	app = app_param
}

func RenderTemplate(w http.ResponseWriter, targetTmpl string) {
	// creer un nouveau TC
	var monCache map[string]*template.Template
	if app.UseCache {
		monCache = app.TemplateCache
	} else {
		monCache, _ = CreateTemplateCache()
	}


	// Obtenir le template demandé à partir du cache
	tmpl, ok := monCache[targetTmpl]

	if !ok {
		log.Fatal("Could not get the template from cache")
	}
	// Avant d'executer le template, tu as besoin de l'executer dans un buffer
	buffer := new(bytes.Buffer)
	err := tmpl.Execute(buffer, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	buffer.WriteTo(w)

}


func CreateTemplateCache() (map[string]*template.Template,error) {

	// Parcours les fichiers de pages ( Go through the pages files )
	// Transforme les pages en *template.Template
	// Parcours les fichiers de layouts ( Go through the layouts files )
	// Make sure the transformed pages are also parsed with the layouts
	// Ecris dans le Writer ( Write to the Writer) <- Later 

	var monCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return monCache, err
	}
	
	for _, page := range pages {
		var fileName string = filepath.Base(page)
		t, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return monCache, err
		}
		
		layouts, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return monCache, err
		}
		
		// It means the layouts exist
		if len(layouts) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return monCache, err
			}
		}

			monCache[fileName] = t
	}

	return monCache, nil

}
