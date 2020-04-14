package cmd

import "github.com/tdekeyser/rite/core/domain"

type Env struct {
	rdb domain.RiteRepository
	tdb domain.TagRepository
}

func NewEnv(rdb domain.RiteRepository, tdb domain.TagRepository) *Env {
	return &Env{rdb: rdb, tdb: tdb}
}
