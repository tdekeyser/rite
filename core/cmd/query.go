package cmd

import "github.com/tdekeyser/rite/core/domain"

func GetRiteQuery(title string, e *Env) *domain.Rite {
	r := e.rdb.Get(title)
	if r == nil {
		return &domain.Rite{Title: title}
	}
	return r
}

func GetAllTitlesQuery(e *Env) []string {
	return e.rdb.GetTitles()
}
