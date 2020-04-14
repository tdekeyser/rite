package filestorage

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain"
	"testing"
)

func TestTagRepository_Create(t *testing.T) {
	tag := domain.Tag("hello")
	r := NewTagRepository(&dataStore{Tags: make(map[domain.Tag]bool)})

	err := r.Create(&tag)
	assert.NoError(t, err)
	assert.True(t, r.DB.Tags["hello"])
}

func TestTagRepository_GetAll(t *testing.T) {
	tags := map[domain.Tag]bool{
		"hello": true,
		"a-tag": true,
	}
	r := NewTagRepository(&dataStore{})
	r.DB.Tags = tags

	actual := r.GetAll()

	assert.Contains(t, actual, domain.Tag("hello"))
	assert.Contains(t, actual, domain.Tag("a-tag"))
	assert.Equal(t, len(actual), 2)
}
