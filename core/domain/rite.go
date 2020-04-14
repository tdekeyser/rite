package domain

import "github.com/google/uuid"

type Rite struct {
	Id    uuid.UUID       `json:"id"`
	Title string          `json:"title"`
	Body  []byte          `json:"body"`
	Tags  map[string]bool `json:"tags"`
}

func NewRite(title string, body string, tags ...string) *Rite {
	tagSet := make(map[string]bool)
	r := &Rite{
		Id:    uuid.New(),
		Title: title,
		Body:  []byte(body),
		Tags:  tagSet,
	}
	for _, t := range tags {
		r.AddTag(t)
	}
	return r
}

func (r *Rite) AddTag(t string) {
	r.Tags[t] = true
}

type RiteRepository interface {

	// Create stores the input Rite such that it can be
	// retrieved again later.
	Create(*Rite) error

	// Get returns a Rite based on a string identifier.
	Get(string) *Rite

	// GetTitles returns all unique identifiers associated
	// with Rites in the repository.
	GetTitles() []string
}
