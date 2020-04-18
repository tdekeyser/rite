package cmd

import (
	"github.com/tdekeyser/rite/core/domain/rite"
)

type Env struct {
	rdb rite.Repository
}

func NewEnv(rdb rite.Repository) *Env {
	return &Env{rdb: rdb}
}
