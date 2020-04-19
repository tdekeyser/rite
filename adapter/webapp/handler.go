package webapp

import (
	"github.com/tdekeyser/rite/core/cmd"
	"log"
	"net/http"
)

const (
	View = "/v/"
	Save = "/s/"
	All  = "/a/"
	Tags = "/t/"
)

func NewHandler(h func(http.ResponseWriter, *http.Request, *cmd.Env), c *cmd.Env) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r, c)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, e *cmd.Env) {
	t := r.URL.Path[len(View):]
	rt := cmd.RiteQuery(t, e)
	renderTemplate(w, Table, rt)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, e *cmd.Env) {
	t := r.URL.Path[len(Save):]
	b := r.FormValue("body")
	tg := r.FormValue("tag")

	if tg != "" {
		err := cmd.AddTagCommand(t, tg, e)
		if err != nil {
			http.NotFound(w, r)
		}
	}

	if b != "" {
		err := cmd.UpdateBodyCommand(t, b, e)
		if err != nil {
			http.NotFound(w, r)
		}
	}

	http.Redirect(w, r, "/v/"+t, http.StatusFound)
}

func TitlesHandler(w http.ResponseWriter, _ *http.Request, e *cmd.Env) {
	ts := cmd.AllRiteTitlesQuery(e)
	renderTemplate(w, Overview, ts)
}

func TagsHandler(w http.ResponseWriter, r *http.Request, e *cmd.Env) {
	t, err := cmd.AllTagsAndSomeTitleQuery(e)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
	}
	renderTemplate(w, OverviewTags, t)
}
