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

func NewHandler(h func(http.ResponseWriter, *http.Request, *cmd.Env), c *cmd.Env) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r, c)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, e *cmd.Env) {
	t := r.URL.Path[len(View):]
	rite := cmd.GetRiteQuery(t, e)
	renderTemplate(w, Table, rite)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, e *cmd.Env) {
	t := r.URL.Path[len(Save):]
	b := r.FormValue("body")
	err := cmd.SaveRiteCommand(t, b, e)
	if err != nil {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, "/v/"+t, http.StatusFound)
	}
}

func TitlesHandler(w http.ResponseWriter, _ *http.Request, e *cmd.Env) {
	ts := cmd.GetAllTitlesQuery(e)
	renderTemplate(w, Overview, ts)
}
