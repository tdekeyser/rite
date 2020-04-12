package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain"
	"testing"
)

func TestModule_GetRite(t *testing.T) {
	r := &domain.Rite{Title: "1", Body: []byte("hello there")}
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("Get", "1").Return(r)

	actual := GetRiteQuery("1", e)

	m.AssertExpectations(t)
	assert.Equal(t, r, actual)
}

func TestModule_GetRite_none_found_returns_emptyRite(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("Get", "100").Return(nil)

	actual := GetRiteQuery("100", e)

	m.AssertExpectations(t)
	assert.Equal(t, &domain.Rite{Title: "100"}, actual)
}
