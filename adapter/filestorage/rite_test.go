package filestorage

import (
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain"
	"testing"
)

const dbTest = "test_db.json"

func TestDb_Save(t *testing.T) {
	conn := RiteRepository{DB: &dataStore{}}
	r := domain.Rite{Title: "1", Body: []byte("hello"), Tags: []string{"a-tag"}}

	err := conn.Create(&r)
	assert.NoError(t, err)

	assert.Contains(t, conn.DB.Rites, r)
}

func TestDb_Save_multiple(t *testing.T) {
	conn := RiteRepository{DB: &dataStore{}}
	r1 := domain.Rite{Title: "1", Body: []byte("hello"), Tags: []string{"a-tag"}}
	r2 := domain.Rite{Title: "2", Body: []byte("hi there")}

	assert.NoError(t, conn.Create(&r1))
	assert.NoError(t, conn.Create(&r2))

	assert.Contains(t, conn.DB.Rites, r1, r2)
}

func TestDb_Get(t *testing.T) {
	r := domain.Rite{Title: "1", Body: []byte("hello")}
	conn := RiteRepository{DB: &dataStore{Rites: []domain.Rite{r}}}

	assert.Equal(t, r, *conn.Get("1"))
}

func TestDb_Get_does_not_return_copy(t *testing.T) {
	r := domain.Rite{Title: "1", Body: []byte("hello")}
	conn := RiteRepository{DB: &dataStore{Rites: []domain.Rite{r}}}
	actual := conn.Get("1")

	actual.Body = []byte(("something else"))

	assert.Equal(t, conn.DB.Rites[0].Body, []byte("something else"))
}

func TestDb_GetAll(t *testing.T) {
	dbName = dbTest
	rs := []domain.Rite{
		{Title: "1", Body: []byte("hello")},
		{Title: "2", Body: []byte("other text")},
	}
	ts := []string{"1", "2"}
	conn := RiteRepository{DB: &dataStore{Rites: rs}}

	assert.Equal(t, ts, conn.GetTitles())
}
