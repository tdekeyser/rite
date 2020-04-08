package domain

type Storage interface {

	// Save stores the input Rite such that it can be
	// retrieved again later.
	// Save overrides the Rite if it exists already.
	Save(*Rite) error

	// Get returns a Rite based on a string identifier.
	Get(string) *Rite
}

type Rite struct {
	Title string `json:"title"`
	Body  []byte `json:"body"`
}
