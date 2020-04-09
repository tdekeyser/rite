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

var version = "0.0"

func main() {
	log.Printf(header, version)

	db, err := filestorage.Open("")
	if err != nil {
		log.Fatal(err)
	}

	m := cmd.NewModule(db)

	http.HandleFunc("/v/", webapp.NewHandler(webapp.ViewHandler, m))
	http.HandleFunc("/s/", webapp.NewHandler(webapp.SaveHandler, m))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
