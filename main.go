package main

import (
	"github.com/tdekeyser/rite/adapter/filestorage"
	"github.com/tdekeyser/rite/adapter/webapp"
	"github.com/tdekeyser/rite/core/cmd"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	deferClose(db)

	r := filestorage.NewRiteRepository(db)
	t := filestorage.NewTagRepository(db)

	e := cmd.NewEnv(r, t)

	http.HandleFunc(webapp.View, webapp.NewHandler(webapp.ViewHandler, e))
	http.HandleFunc(webapp.Save, webapp.NewHandler(webapp.SaveHandler, e))
	http.HandleFunc(webapp.All, webapp.NewHandler(webapp.TitlesHandler, e))
	http.HandleFunc(webapp.Tags, webapp.NewHandler(webapp.TagsHandler, e))

	http.Handle("/assets/", webapp.AssetHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func deferClose(c io.Closer) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-s
		err := c.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
		os.Exit(0)
	}()
}
