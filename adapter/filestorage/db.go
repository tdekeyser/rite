// Package filedb is a way of quickly saving rites to json format.
// It is absolutely NOT thread safe, but that's fine for
// the current use case.
package filestorage

import (
	"encoding/json"
	"github.com/tdekeyser/rite/core/domain/rite"
	"io/ioutil"
	"log"
)

var dbName = "rite_filedb.json"

type dataStore struct {
	Loc   string                       `json:"location"`
	Rites []rite.Rite                  `json:"rites"`
	Tags  map[rite.Tag]map[string]bool `json:"tags"`
}

func newDataStore(loc string) *dataStore {
	return &dataStore{
		Loc:  loc,
		Tags: make(map[rite.Tag]map[string]bool),
	}
}

func Open(location string) (*dataStore, error) {
	conn, err := openExisting(location)
	if err != nil {
		log.Print("Creating new database.")
		err := ioutil.WriteFile(location+dbName, []byte{}, 0600)
		if err != nil {
			panic("Error initiating database: " + err.Error())
		}
		return newDataStore(location), nil
	}
	log.Printf("Found existing database with %v rite(s).", len(conn.Rites))
	return conn, err
}

func (ds *dataStore) Close() error {
	return ds.saveToDisk()
}

func openExisting(loc string) (*dataStore, error) {
	f, err := ioutil.ReadFile(loc + dbName)
	if err != nil {
		return nil, err
	}

	var d dataStore
	err = json.Unmarshal(f, &d)
	if err != nil {
		log.Printf("Could not unmarshal: %v", err)
		return nil, err
	}
	return &d, nil
}

func (ds *dataStore) saveToDisk() error {
	d, err := json.Marshal(ds)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ds.Loc+dbName, d, 0600)
}
