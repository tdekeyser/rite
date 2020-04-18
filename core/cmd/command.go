package cmd

import (
	"github.com/tdekeyser/rite/core/domain/rite"
)

func UpdateBodyCommand(title string, body string, e *Env) error {
	if r, ok := e.rdb.GetByTitle(title); ok {
		r.Body = []byte(body)
		return e.rdb.Update(&r)
	}
	return e.rdb.Create(rite.New(title, body))
}

func AddTagCommand(title string, tag string, e *Env) error {
	if r, ok := e.rdb.GetByTitle(title); ok {
		r.AddTag(rite.Tag(tag))
		return e.rdb.Update(&r)
	}
	return e.rdb.Create(rite.New(title, "", tag))
}
