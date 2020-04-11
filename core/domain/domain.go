package domain

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

type Rite struct {
	Title string `json:"title"`
	Body  []byte `json:"body"`
}
