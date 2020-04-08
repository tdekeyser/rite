package template

import (
	"github.com/tdekeyser/rite/domain"
	"html/template"
	"net/http"
)

const (
	dir   = "template/"
	Table = "table.html"
)

var templates = template.Must(template.ParseFiles(dir + Table))

func Render(w http.ResponseWriter, tmpl string, r *domain.Rite) {
	err := templates.ExecuteTemplate(w, tmpl, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
