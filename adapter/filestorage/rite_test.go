package filestorage

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain/rite"
	"testing"
)

func TestRiteRepository_Create(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	r := rite.New("1", "hello", "a-tag")

	assert.NoError(t, conn.Create(r))

	assert.Contains(t, conn.DB.Rites, *r)
	assert.Contains(t, conn.DB.Tags["a-tag"], "1")
}

func TestRiteRepository_Create_multiple(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	r1 := rite.New("1", "hello", "a-tag")
	r2 := rite.New("2", "hi there", "a-tag")

	assert.NoError(t, conn.Create(r1))
	assert.NoError(t, conn.Create(r2))

	assert.Contains(t, conn.DB.Rites, *r1, *r2)
	assert.Contains(t, conn.DB.Tags["a-tag"], "1")
	assert.Contains(t, conn.DB.Tags["a-tag"], "2")
}

func TestRiteRepository_Update(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	r1 := rite.New("1", "hello", "paper")
	r2 := rite.New("1", "hi there", "paper", "books")
	assert.NoError(t, conn.Create(r1))

	assert.NoError(t, conn.Update(r2))

	assert.Contains(t, conn.DB.Rites, *r2)
	assert.NotContains(t, conn.DB.Rites, *r1)
	assert.Contains(t, conn.DB.Tags["paper"], "1")
	assert.Contains(t, conn.DB.Tags["books"], "1")
}

func TestRiteRepository_Update_removeTags(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	r1 := rite.New("1", "hello", "paper", "books")
	r2 := rite.New("1", "hi there", "paper")
	assert.NoError(t, conn.Create(r1))

	assert.NoError(t, conn.Update(r2))

	assert.Contains(t, conn.DB.Rites, *r2)
	assert.NotContains(t, conn.DB.Rites, *r1)
	assert.Contains(t, conn.DB.Tags["paper"], "1")
	assert.Empty(t, conn.DB.Tags["books"])
}

func TestRiteRepository_GetByTitle(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	r := rite.New("1", "", "books")
	assert.NoError(t, conn.Create(r))

	actual, _ := conn.GetByTitle("1")
	assert.Equal(t, *r, actual)
}

func TestRiteRepository_GetTitles(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	assert.NoError(t, conn.Create(rite.New("1", "", "books")))
	assert.NoError(t, conn.Create(rite.New("2", "", "paper")))

	assert.Equal(t, []string{"1", "2"}, conn.GetTitles())
}

func TestRiteRepository_GetTitlesByTag(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	assert.NoError(t, conn.Create(rite.New("1", "", "books")))
	assert.NoError(t, conn.Create(rite.New("2", "", "paper")))
	assert.NoError(t, conn.Create(rite.New("3", "", "books", "paper")))

	assert.Contains(t, conn.GetTitlesByTag("books"), "1")
	assert.Contains(t, conn.GetTitlesByTag("books"), "3")
	assert.Contains(t, conn.GetTitlesByTag("paper"), "2")
	assert.Contains(t, conn.GetTitlesByTag("paper"), "3")
	assert.Empty(t, conn.GetTitlesByTag("non-existent"))
}

func TestRiteRepository_GetTags(t *testing.T) {
	conn := RiteRepository{DB: newDataStore("")}
	assert.NoError(t, conn.Create(rite.New("1", "", "books")))
	assert.NoError(t, conn.Create(rite.New("2", "", "paper")))
	assert.NoError(t, conn.Create(rite.New("3", "", "books", "paper")))

	assert.Contains(t, conn.GetTags(), rite.Tag("books"), rite.Tag("paper"))
}
