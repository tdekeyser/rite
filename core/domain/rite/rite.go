package rite

import "github.com/google/uuid"

type Tag string

type Rite struct {
	Id    uuid.UUID    `json:"id"`
	Title string       `json:"title"`
	Body  []byte       `json:"body"`
	Tags  map[Tag]bool `json:"tags"`
}

func New(title string, body string, tags ...string) *Rite {
	tagSet := make(map[Tag]bool)
	for _, t := range tags {
		tagSet[Tag(t)] = true
	}
	return &Rite{
		Id:    uuid.New(),
		Title: title,
		Body:  []byte(body),
		Tags:  tagSet,
	}
}

func (r *Rite) Copy() Rite {
	// A copy of a map is always a copy of the reference!
	// In order to make a copy of the Rite, we need to
	// make an explicit copy of its tags.
	tagsCopy := make(map[Tag]bool)
	for t := range r.Tags {
		tagsCopy[t] = true
	}
	return Rite{
		Id:    r.Id,
		Title: r.Title,
		Body:  r.Body,
		Tags:  tagsCopy,
	}
}

func (r *Rite) AddTag(t Tag) {
	r.Tags[t] = true
}

type Repository interface {
	Create(*Rite) error
	Update(*Rite) error
	GetByTitle(string) (Rite, bool)
	GetTitles() []string
	GetTitlesByTag(Tag) []string
	GetTags() []Tag
}
