package cmd

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/tdekeyser/rite/core/domain/rite"
)

type RiteRepositoryMock struct {
	mock.Mock
}

func (db *RiteRepositoryMock) Create(r *rite.Rite) error {
	r.Id = uuid.Nil
	v := db.Called(r)
	return v.Error(0)
}

func (db *RiteRepositoryMock) Update(r *rite.Rite) error {
	v := db.Called(r)
	return v.Error(0)
}

func (db *RiteRepositoryMock) GetByTitle(title string) (rite.Rite, bool) {
	v := db.Called(title)
	return v.Get(0).(rite.Rite), v.Get(1).(bool)
}

func (db *RiteRepositoryMock) GetTitlesByTag(t rite.Tag) []string {
	v := db.Called(t)
	return v.Get(0).([]string)
}

func (db *RiteRepositoryMock) GetTitles() []string {
	v := db.Called()
	return v.Get(0).([]string)
}

func (db *RiteRepositoryMock) GetTags() []rite.Tag {
	v := db.Called()
	return v.Get(0).([]rite.Tag)
}
