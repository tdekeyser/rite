package webapp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/tdekeyser/rite/core/domain/rite"
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

var templates = template.Must(
	template.New("").Funcs(template.FuncMap{
		"str2col": str2col,
	}).ParseFiles(
		appDir+"templates/"+header,
		appDir+"templates/"+footer,
		appDir+"templates/"+Table,
		appDir+"templates/"+Overview,
		appDir+"templates/"+OverviewTags,
	))

func str2col(s rite.Tag) string {
	hash := md5.Sum([]byte(s))
	return fmt.Sprintf("#%s", hex.EncodeToString(hash[:])[:6])
}

func renderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
