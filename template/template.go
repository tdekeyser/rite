package template

import (
	"github.com/tdekeyser/rite/domain"
	"html/template"
	"net/http"
	"os"
)

const Table = "table.html"

var templateDir = os.Getenv("RITE_TEMPLATE_DIR")

var templates = template.Must(template.ParseFiles(templateDir + Table))

func Render(w http.ResponseWriter, tmpl string, r *domain.Rite) {
	err := templates.ExecuteTemplate(w, tmpl, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
