package cmd

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain/rite"
	"testing"
)

func TestSaveRiteCommand(t *testing.T) {
	r := rite.New("1", "hello there")
	r.Id = uuid.Nil
	m := new(RiteRepositoryMock)
	m.On("GetByTitle", r.Title).Return(rite.Rite{}, false)
	m.On("Create", r).Return(nil)
	e := NewEnv(m)

	err := UpdateBodyCommand("1", "hello there", e)
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestGetAllTitlesQuery(t *testing.T) {
	ts := []string{"1", "2"}
	m := new(RiteRepositoryMock)
	m.On("GetTitles").Return(ts)
	e := NewEnv(m)

	assert.Equal(t, ts, AllRiteTitlesQuery(e))
	m.AssertExpectations(t)
}

func TestAddTagCommand(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)
	r := rite.New("1", "hello there", "a-tag")
	m.On("GetByTitle", "1").Return(*r, true)
	m.On("Update", r).Return(nil)

	err := AddTagCommand("1", "a-tag", e)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}
