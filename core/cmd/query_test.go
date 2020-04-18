package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain/rite"
	"testing"
)

func TestRiteQuery(t *testing.T) {
	r := &rite.Rite{Title: "1", Body: []byte("hello there")}
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("GetByTitle", "1").Return(*r, true)

	actual := RiteQuery("1", e)

	m.AssertExpectations(t)
	assert.Equal(t, r, actual)
}

func TestRiteQuery_none_found_returns_emptyRite(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("GetByTitle", "100").Return(rite.Rite{}, false)

	actual := RiteQuery("100", e)

	m.AssertExpectations(t)
	assert.Equal(t, &rite.Rite{Title: "100"}, actual)
}
