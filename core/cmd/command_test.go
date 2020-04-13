package cmd

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain"
	"testing"
)

func TestSaveRiteCommand(t *testing.T) {
	r := domain.NewRite("1", "hello there")
	r.Id = uuid.Nil
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("Get", r.Title).Return(nil)
	m.On("Create", r).Return(nil)

	err := UpdateBodyCommand("1", "hello there", e)
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestSaveRiteCommand_updatesExisting(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	r := domain.NewRite("1", "hello there")
	rUpdated := r
	rUpdated.Body = []byte("other text")

	m.On("Get", "1").Return(r)

	err := UpdateBodyCommand("1", "other text", e)
	assert.NoError(t, err)

	m.AssertExpectations(t)
	assert.Equal(t, r, rUpdated)
}

func TestGetAllTitlesQuery(t *testing.T) {
	ts := []string{"1", "2"}
	m := new(RiteRepositoryMock)
	e := NewEnv(m)

	m.On("GetTitles").Return(ts)

	assert.Equal(t, ts, GetAllTitlesQuery(e))
	m.AssertExpectations(t)
}

func TestAddTagCommand(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)
	r := domain.NewRite("1", "hello there", "a-tag")

	m.On("Get", "1").Return(r)

	err := AddTagCommand("1", "a-tag", e)
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestAddTagCommand_updatesRite(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)
	r := domain.NewRite("1", "hello there")

	m.On("Get", "1").Return(r)

	err := AddTagCommand("1", "a-tag", e)
	assert.NoError(t, err)
	m.AssertExpectations(t)

	assert.True(t, r.Tags["a-tag"])
}

func TestAddTagCommand_sameTag_notAddedTwice(t *testing.T) {
	m := new(RiteRepositoryMock)
	e := NewEnv(m)
	r := domain.NewRite("1", "hello there", "a-tag")

	m.On("Get", "1").Return(r)

	err := AddTagCommand("1", "a-tag", e)

	assert.NoError(t, err)
	m.AssertExpectations(t)

	assert.True(t, r.Tags["a-tag"])
}
