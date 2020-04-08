package main

import (
	"github.com/tdekeyser/rite/domain"
	"github.com/tdekeyser/rite/filestorage"
	"html/template"
	"log"
	"net/http"
)

const header = `
╋╋┏┳┓
┏┳╋┫┗┳━┓
┃┏┫┃┏┫┻┫ 
┗┛┗┻━┻━┛   v%v
`

var (
	version   = "0.0"
	templates = template.Must(template.ParseFiles("table.html"))
	db        domain.Storage
)

func renderTemplate(w http.ResponseWriter, tmpl string, r *domain.Rite) {
	err := templates.ExecuteTemplate(w, tmpl, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/v/"):]
	rite := db.Get(title)
	if rite == nil {
		rite = &domain.Rite{Title: title}
	}
	renderTemplate(w, "table.html", rite)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/s/"):]
	body := r.FormValue("body")
	rite := &domain.Rite{Title: title, Body: []byte(body)}
	err := db.Save(rite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/v/"+title, http.StatusFound)
}

func main() {
	log.Printf(header, version)

	var err error
	db, err = filestorage.Open("")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/v/", viewHandler)
	http.HandleFunc("/s/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
