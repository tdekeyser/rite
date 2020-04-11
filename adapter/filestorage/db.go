// Package filedb is a way of quickly saving data to json format.
// It is absolutely NOT thread safe, but that's fine for
// the current use case.
package filestorage

import (
	"encoding/json"
	"github.com/tdekeyser/rite/core/domain"
	"io/ioutil"
	"log"
)

var dbName = "rite_filedb.json"

type db struct {
	loc  string
	data []domain.Rite
}

func Open(location string) (*db, error) {
	conn, err := openExisting(location)
	if err != nil {
		log.Print("Creating new database.")
		return newDb(location), nil
	}
	log.Printf("Found existing database with %v rite(s).", len(conn.data))
	return conn, err
}

func newDb(loc string) *db {
	err := ioutil.WriteFile(loc+dbName, []byte{}, 0600)
	if err != nil {
		panic("Error initiating database: " + err.Error())
	}
	return &db{loc: loc}
}

func openExisting(loc string) (*db, error) {
	f, err := ioutil.ReadFile(loc + dbName)
	if err != nil {
		return nil, err
	}

	var data []domain.Rite
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Printf("Could not unmarshal data: %v", err)
		return nil, err
	}

	return &db{loc, data}, nil
}

func (conn *db) Save(r *domain.Rite) error {
	prev := conn.Get(r.Title)
	if prev != nil {
		prev.Body = r.Body
	} else {
		conn.data = append(conn.data, *r)
	}
	err := conn.saveToDisk()
	if err != nil {
		log.Printf("Could not save data to disk: %v", err)
	}
	return err
}

func (conn *db) saveToDisk() error {
	data, err := json.Marshal(conn.data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(conn.loc+dbName, data, 0600)
}

func (conn *db) Get(title string) *domain.Rite {
	for i, r := range conn.data {
		if title == r.Title {
			return &conn.data[i]
		}
	}
	return nil
}

func (conn *db) GetIds() []string {
	var ts []string
	for _, r := range conn.data {
		ts = append(ts, r.Title)
	}
	return ts
}
