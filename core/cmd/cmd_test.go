package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tdekeyser/rite/core/domain"
	"testing"
)

func TestModule_GetRite(t *testing.T) {
	r := &domain.Rite{Title: "1", Body: []byte("hello there")}
	m := new(RiteRepositoryMock)
	c := Module{db: m}

	m.On("Get", "1").Return(r)

	actual := c.GetRite("1")

	m.AssertExpectations(t)
	assert.Equal(t, r, actual)
}

func TestModule_GetRite_none_found_returns_emptyRite(t *testing.T) {
	m := new(RiteRepositoryMock)
	c := Module{db: m}

	m.On("Get", "100").Return(nil)

	actual := c.GetRite("100")

	m.AssertExpectations(t)
	assert.Equal(t, &domain.Rite{Title: "100"}, actual)
}

func TestModule_SaveRite(t *testing.T) {
	r := &domain.Rite{Title: "1", Body: []byte("hello there")}
	m := new(RiteRepositoryMock)
	c := Module{db: m}

	m.On("Save", r).Return(nil)

	err := c.SaveRite("1", "hello there")
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

type RiteRepositoryMock struct {
	mock.Mock
}

func (db *RiteRepositoryMock) Save(r *domain.Rite) error {
	v := db.Called(r)
	return v.Error(0)
}

func (db *RiteRepositoryMock) Get(title string) *domain.Rite {
	v := db.Called(title)
	r := v.Get(0)
	if r != nil {
		return r.(*domain.Rite)
	}
	return nil
}
