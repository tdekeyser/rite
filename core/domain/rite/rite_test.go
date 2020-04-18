package rite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRite_AddTag_TagsAreASet(t *testing.T) {
	r := New("1", "hello there", "a-tag")

	r.AddTag("a-tag")

	assert.True(t, r.Tags["a-tag"])
}
