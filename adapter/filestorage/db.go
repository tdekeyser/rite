// Package filedb is a way of quickly saving rites to json format.
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

type dataStore struct {
	loc   string
	rites []domain.Rite
}

func Open(location string) (*dataStore, error) {
	conn, err := openExisting(location)
	if err != nil {
		log.Print("Creating new database.")
		return newDb(location), nil
	}
	log.Printf("Found existing database with %v rite(s).", len(conn.rites))
	return conn, err
}

func newDb(loc string) *dataStore {
	err := ioutil.WriteFile(loc+dbName, []byte{}, 0600)
	if err != nil {
		panic("Error initiating database: " + err.Error())
	}
	return &dataStore{loc: loc}
}

func openExisting(loc string) (*dataStore, error) {
	f, err := ioutil.ReadFile(loc + dbName)
	if err != nil {
		return nil, err
	}

	var data []domain.Rite
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Printf("Could not unmarshal rites: %v", err)
		return nil, err
	}

	return &dataStore{loc, data}, nil
}

func (conn *dataStore) saveToDisk() error {
	data, err := json.Marshal(conn.rites)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(conn.loc+dbName, data, 0600)
}
