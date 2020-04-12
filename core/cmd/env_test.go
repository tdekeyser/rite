package cmd

import (
	"github.com/stretchr/testify/mock"
	"github.com/tdekeyser/rite/core/domain"
)

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

func (db *RiteRepositoryMock) GetIds() []string {
	v := db.Called()
	return v.Get(0).([]string)
}