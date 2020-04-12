package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain"
	"testing"
)

func TestModule_SaveRite(t *testing.T) {
	r := &domain.Rite{Title: "1", Body: []byte("hello there")}
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("Save", r).Return(nil)

	err := SaveRiteCommand("1", "hello there", e)
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestModule_GetAllRites(t *testing.T) {
	ts := []string{"1", "2"}
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("GetIds").Return(ts)

	assert.Equal(t, ts, GetAllTitlesQuery(e))
	m.AssertExpectations(t)
}

func TestAddTagCommand(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)
	r := domain.NewRite("1", "hello there")
	rWithTag := r
	rWithTag.Tags = []string{"a-tag"}

	m.On("Get", "1").Return(r)
	m.On("Save", rWithTag).Return(nil)

	err := AddTagCommand("1", "a-tag", e)
	assert.NoError(t, err)
	m.AssertExpectations(t)
}
