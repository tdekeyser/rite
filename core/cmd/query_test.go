package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain/rite"
	"testing"
)

func TestRiteQuery(t *testing.T) {
	r := &rite.Rite{Title: "1", Body: []byte("hello there")}
	m := new(RiteRepositoryMock)
	m.On("GetByTitle", "1").Return(*r, true)
	e := NewEnv(m)

	actual := RiteQuery("1", e)

	m.AssertExpectations(t)
	assert.Equal(t, *r, actual)
}

func TestRiteQuery_none_found_returns_emptyRite(t *testing.T) {
	m := new(RiteRepositoryMock)
	m.On("GetByTitle", "100").Return(rite.Rite{}, false)
	e := NewEnv(m)

	actual := RiteQuery("100", e)

	m.AssertExpectations(t)
	assert.Equal(t, rite.Rite{Title: "100"}, actual)
}

func TestAllTagsAndSomeTitleQuery(t *testing.T) {
	m := new(RiteRepositoryMock)
	m.On("GetTags").Return([]rite.Tag{"books", "paper"})
	m.On("GetTitlesByTag", rite.Tag("books")).Return([]string{"1", "2"})
	m.On("GetTitlesByTag", rite.Tag("paper")).Return([]string{"2", "3"})
	e := NewEnv(m)

	actual, err := AllTagsAndSomeTitleQuery(e)

	m.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, rite.Tag("books"), actual[0].Tag)
	assert.True(t, actual[0].Title == "1" || actual[0].Title == "2")
	assert.Equal(t, rite.Tag("paper"), actual[1].Tag)
	assert.True(t, actual[1].Title == "2" || actual[1].Title == "3")
}
