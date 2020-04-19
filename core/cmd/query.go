package cmd

import (
	"fmt"
	"github.com/tdekeyser/rite/core/domain/rite"
	"math/rand"
)

func RiteQuery(title string, e *Env) rite.Rite {
	if r, ok := e.rdb.GetByTitle(title); ok {
		return r
	}
	return rite.Rite{Title: title}
}

func AllRiteTitlesQuery(e *Env) []string {
	return e.rdb.GetTitles()
}

type TagTitle struct {
	Tag   rite.Tag
	Title string
}

func AllTagsAndSomeTitleQuery(e *Env) ([]TagTitle, error) {
	tags := e.rdb.GetTags()
	tt := make([]TagTitle, len(tags))
	for i, t := range tags {
		r := e.rdb.GetTitlesByTag(t)
		if len(r) == 0 {
			return nil, fmt.Errorf("no rites found for tag %s", t)
		}
		tt[i] = TagTitle{t, r[rand.Intn(len(r))]}
	}
	return tt, nil
}
