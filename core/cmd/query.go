package cmd

import "github.com/tdekeyser/rite/core/domain"

func RiteQuery(title string, e *Env) *domain.Rite {
	r := e.rdb.Get(title)
	if r == nil {
		return &domain.Rite{Title: title}
	}
	return r
}

func AllRiteTitlesQuery(e *Env) []string {
	return e.rdb.GetTitles()
}
