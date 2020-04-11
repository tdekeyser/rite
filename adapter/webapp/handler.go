package webapp

import (
	"github.com/tdekeyser/rite/core/cmd"
	"net/http"
)

const (
	View = "/v/"
	Save = "/s/"
	All  = "/a/"
)

func NewHandler(h func(http.ResponseWriter, *http.Request, *cmd.Module), c *cmd.Module) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r, c)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, m *cmd.Module) {
	t := r.URL.Path[len(View):]

	rite := m.GetRite(t)

	renderTemplate(w, Table, rite)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, m *cmd.Module) {
	t := r.URL.Path[len(Save):]
	b := r.FormValue("body")

	err := m.SaveRite(t, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/v/"+t, http.StatusFound)
}

func TitlesHandler(w http.ResponseWriter, _ *http.Request, m *cmd.Module) {
	ts := m.GetAllRiteTitles()
	renderTemplate(w, Overview, ts)
}
