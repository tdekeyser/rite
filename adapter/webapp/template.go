package webapp

import (
	"github.com/tdekeyser/rite/core/domain"
	"html/template"
	"net/http"
	"os"
)

const Table = "table.html"

var appDir = os.Getenv("RITE_APP_DIR")

var AssetHandler = http.StripPrefix("/assets/", http.FileServer(http.Dir(appDir+"assets/")))
var templates = template.Must(template.ParseFiles(appDir + "templates/" + Table))

func renderTemplate(w http.ResponseWriter, tmpl string, r *domain.Rite) {
	err := templates.ExecuteTemplate(w, tmpl, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
