package handlers

import (
	"net/http"

	"github.com/sonu-warsi/bookings/pkg/config"
	"github.com/sonu-warsi/bookings/pkg/models"
	"github.com/sonu-warsi/bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository type
type Repository struct {
	App *config.AppConfig
}

//Create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//Sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

//home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["app"] = "Hello, world!"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
