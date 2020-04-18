package filestorage

import (
	"github.com/tdekeyser/rite/core/domain/rite"
)

type RiteRepository struct {
	DB *dataStore
}

func NewRiteRepository(db *dataStore) *RiteRepository {
	return &RiteRepository{DB: db}
}

func (repo *RiteRepository) Create(r *rite.Rite) error {
	repo.DB.Rites = append(repo.DB.Rites, *r)
	for t := range r.Tags {
		repo.addTag(t, r.Title)
	}
	return nil
}

func (repo *RiteRepository) addTag(t rite.Tag, title string) {
	if repo.DB.Tags[t] == nil {
		repo.DB.Tags[t] = make(map[string]bool)
	}
	repo.DB.Tags[t][title] = true
}

func (repo *RiteRepository) Update(r *rite.Rite) error {
	var oldTags map[rite.Tag]bool

	for i, old := range repo.DB.Rites {
		if r.Title == old.Title {
			oldTags = old.Tags
			repo.DB.Rites[i] = *r
		}
	}

	applyToDifference(oldTags, r.Tags, func(t rite.Tag) {
		delete(repo.DB.Tags[t], r.Title)
	})
	applyToDifference(r.Tags, oldTags, func(t rite.Tag) {
		repo.addTag(t, r.Title)
	})

	return nil
}

func applyToDifference(a map[rite.Tag]bool, b map[rite.Tag]bool, f func(rite.Tag)) {
	for t := range a {
		if _, ok := b[t]; !ok {
			f(t)
		}
	}
}

func (repo *RiteRepository) GetByTitle(title string) (rite.Rite, bool) {
	for _, r := range repo.DB.Rites {
		if title == r.Title {
			return r.Copy(), true
		}
	}
	return rite.Rite{}, false
}

func (repo *RiteRepository) GetTitlesByTag(t rite.Tag) []string {
	var ts []string
	for t := range repo.DB.Tags[t] {
		ts = append(ts, t)
	}
	return ts
}

func (repo *RiteRepository) GetTitles() []string {
	var ts []string
	for _, r := range repo.DB.Rites {
		ts = append(ts, r.Title)
	}
	return ts
}

func (repo *RiteRepository) GetTags() []rite.Tag {
	var ts []rite.Tag
	for t := range repo.DB.Tags {
		ts = append(ts, t)
	}
	return ts
}
