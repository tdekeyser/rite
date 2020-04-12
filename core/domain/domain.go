package domain

import "github.com/google/uuid"

type Rite struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Body  []byte    `json:"body"`
}

func NewRite(title string, body string) *Rite {
	return &Rite{uuid.New(), title, []byte(body)}
}

type RiteRepository interface {

	// Save stores the input Rite such that it can be
	// retrieved again later.
	// Save overrides the Rite if it exists already.
	Save(*Rite) error

	// Get returns a Rite based on a string identifier.
	Get(string) *Rite

	// GetIds returns all unique identifiers associated
	// with Rites in the repository.
	GetIds() []string
}
