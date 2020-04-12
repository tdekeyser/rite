package filestorage

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tdekeyser/rite/core/domain"
	"io/ioutil"
	"os"
	"testing"
)

const dbTest = "test_db.json"

func TestDb_Save(t *testing.T) {
	dbName = dbTest
	conn := RiteRepository{DB: &dataStore{}}
	r := domain.Rite{Title: "1", Body: []byte("hello")}

	err := conn.Save(&r)
	assert.NoError(t, err)

	assertDbContainsExactly(t, r)

	assert.NoError(t, os.Remove(dbTest))
}

func TestDb_Save_multiple(t *testing.T) {
	dbName = dbTest
	conn := RiteRepository{DB: &dataStore{}}
	r1 := domain.Rite{Title: "1", Body: []byte("hello")}
	r2 := domain.Rite{Title: "2", Body: []byte("hi there")}

	assert.NoError(t, conn.Save(&r1))
	assert.NoError(t, conn.Save(&r2))

	assertDbContainsExactly(t, r1, r2)

	assert.NoError(t, os.Remove(dbTest))
}

func TestDb_Save_overrides_same_title(t *testing.T) {
	dbName = dbTest
	conn := RiteRepository{DB: &dataStore{}}
	r1 := domain.Rite{Title: "1", Body: []byte("hello")}
	r2 := domain.Rite{Title: "1", Body: []byte("other text")}

	assert.NoError(t, conn.Save(&r1))
	assert.NoError(t, conn.Save(&r2))

	assertDbContainsExactly(t, r2)

	assert.NoError(t, os.Remove(dbTest))
}

func assertDbContainsExactly(t *testing.T, r ...domain.Rite) {
	f, err := ioutil.ReadFile(dbTest)
	assert.NoError(t, err)

	var actual []domain.Rite
	err = json.Unmarshal(f, &actual)
	assert.NoError(t, err)

	assert.Equal(t, r, actual)
}

func TestDb_Get(t *testing.T) {
	r := domain.Rite{Title: "1", Body: []byte("hello")}
	conn := RiteRepository{DB: &dataStore{rites: []domain.Rite{r}}}

	assert.Equal(t, r, *conn.Get("1"))
}

func TestDb_Get_does_not_return_copy(t *testing.T) {
	r := domain.Rite{Title: "1", Body: []byte("hello")}
	conn := RiteRepository{DB: &dataStore{rites: []domain.Rite{r}}}
	actual := conn.Get("1")

	actual.Body = []byte(("something else"))

	assert.Equal(t, conn.DB.rites[0].Body, []byte("something else"))
}

func TestDb_GetAll(t *testing.T) {
	dbName = dbTest
	rs := []domain.Rite{
		{Title: "1", Body: []byte("hello")},
		{Title: "2", Body: []byte("other text")},
	}
	ts := []string{"1", "2"}
	conn := RiteRepository{DB: &dataStore{rites: rs}}

	assert.Equal(t, ts, conn.GetIds())
}
