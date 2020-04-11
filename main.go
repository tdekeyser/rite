package main

import (
	"github.com/tdekeyser/rite/adapter/filestorage"
	"github.com/tdekeyser/rite/adapter/webapp"
	"github.com/tdekeyser/rite/core/cmd"
	"log"
	"net/http"
)

const header = `
╋╋┏┳┓
┏┳╋┫┗┳━┓
┃┏┫┃┏┫┻┫ 
┗┛┗┻━┻━┛   v%v
`
const version = "0.0"

var commits = ".0"

func main() {
	log.Printf(header, version+commits)

	db, err := filestorage.Open("")
	if err != nil {
		log.Fatal(err)
	}

	m := cmd.NewModule(db)

	http.HandleFunc(webapp.View, webapp.NewHandler(webapp.ViewHandler, m))
	http.HandleFunc(webapp.Save, webapp.NewHandler(webapp.SaveHandler, m))
	http.HandleFunc(webapp.All, webapp.NewHandler(webapp.TitlesHandler, m))

	http.Handle("/assets/", webapp.AssetHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
