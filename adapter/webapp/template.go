package webapp

import (
	"html/template"
	"net/http"
	"os"
)

const (
	header       = "header.html"
	footer       = "footer.html"
	Table        = "table.html"
	Overview     = "overview.html"
	OverviewTags = "overview_tags.html"
)

var appDir = os.Getenv("RITE_APP_DIR")

var AssetHandler = http.StripPrefix("/assets/", http.FileServer(http.Dir(appDir+"assets/")))

var templates = template.Must(template.ParseFiles(
	appDir+"templates/"+header,
	appDir+"templates/"+footer,
	appDir+"templates/"+Table,
	appDir+"templates/"+Overview,
	appDir+"templates/"+OverviewTags,
))

func renderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
