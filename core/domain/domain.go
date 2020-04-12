package domain

import "github.com/google/uuid"

type Tag string

type Rite struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Body  []byte    `json:"body"`
	Tags  []string  `json:"tags"`
}

func NewRite(title string, body string, tags ...string) *Rite {
	return &Rite{Id: uuid.New(), Title: title, Body: []byte(body), Tags: tags}
}

type RiteRepository interface {

	// Create stores the input Rite such that it can be
	// retrieved again later.
	Create(*Rite) error

	// Get returns a Rite based on a string identifier.
	Get(string) *Rite

	// GetIds returns all unique identifiers associated
	// with Rites in the repository.
	GetIds() []string
}
