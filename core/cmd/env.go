package cmd

import "github.com/tdekeyser/rite/core/domain"

type Env struct {
	rdb domain.RiteRepository
}

func NewEnv(db domain.RiteRepository) *Env {
	return &Env{rdb: db}
}
