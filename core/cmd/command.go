package cmd

import "github.com/tdekeyser/rite/core/domain"

func UpdateBodyCommand(title string, body string, e *Env) error {
	r := e.rdb.Get(title)

	if r == nil {
		return e.rdb.Create(domain.NewRite(title, body))
	}

	r.Body = []byte(body)
	return nil
}

func AddTagCommand(title string, tag string, e *Env) error {
	r := e.rdb.Get(title)

	if r == nil {
		return e.rdb.Create(domain.NewRite(title, "", tag))
	}

	r.AddTag(tag)
	return nil
}
