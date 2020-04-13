package filestorage

import (
	"github.com/tdekeyser/rite/core/domain"
)

type RiteRepository struct {
	DB *dataStore
}

func NewRiteRepository(db *dataStore) *RiteRepository {
	return &RiteRepository{DB: db}
}

func (repo *RiteRepository) Create(r *domain.Rite) error {
	repo.DB.Rites = append(repo.DB.Rites, *r)
	return nil
}

func (repo *RiteRepository) Get(title string) *domain.Rite {
	for i, r := range repo.DB.Rites {
		if title == r.Title {
			return &repo.DB.Rites[i]
		}
	}
	return nil
}

func (repo *RiteRepository) GetTitles() []string {
	var ts []string
	for _, r := range repo.DB.Rites {
		ts = append(ts, r.Title)
	}
	return ts
}
