package cmd

import "github.com/tdekeyser/rite/core/domain"

func SaveRiteCommand(title string, body string, e *Env) error {
	r := domain.NewRite(title, body)
	return e.rdb.Save(r)
}

func AddTagCommand(title string, tag string, e *Env) error {
	r := e.rdb.Get(title)
	r.Tags = append(r.Tags, tag)
	return e.rdb.Save(r)
}
