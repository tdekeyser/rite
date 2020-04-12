package cmd

import "github.com/tdekeyser/rite/core/domain"

func SaveRiteCommand(title string, body string, e *Env) error {
	r := domain.NewRite(title, body)
	return e.rdb.Save(r)
}
