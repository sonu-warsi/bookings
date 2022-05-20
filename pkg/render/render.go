package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sonu-warsi/bookings/pkg/config"
	"github.com/sonu-warsi/bookings/pkg/models"
)

var app *config.AppConfig

//sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buff := new(bytes.Buffer)

	td = AddDefultData(td)

	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser: ", err)
	}

}

//CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	Cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return Cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(template.FuncMap{}).ParseFiles(page)
		if err != nil {
			return Cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return Cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
		}
		if err != nil {
			return Cache, err
		}

		Cache[name] = ts
	}

	return Cache, nil
}
