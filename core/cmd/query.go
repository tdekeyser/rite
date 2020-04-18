package cmd

import (
	"github.com/tdekeyser/rite/core/domain/rite"
)

func RiteQuery(title string, e *Env) *rite.Rite {
	if r, ok := e.rdb.GetByTitle(title); ok {
		return &r
	}
	return &rite.Rite{Title: title}
}

func AllRiteTitlesQuery(e *Env) []string {
	return e.rdb.GetTitles()
}

func AllTagsQuery(e *Env) []rite.Tag {
	return e.rdb.GetTags()
}
