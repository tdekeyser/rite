package filestorage

import (
	"github.com/tdekeyser/rite/core/domain"
	"log"
)

type RiteRepository struct {
	DB *dataStore
}

func NewRiteRepository(db *dataStore) *RiteRepository {
	return &RiteRepository{DB: db}
}

func (repo *RiteRepository) Save(r *domain.Rite) error {
	prev := repo.Get(r.Title)
	if prev != nil {
		prev.Body = r.Body
		prev.Tags = r.Tags
	} else {
		repo.DB.Rites = append(repo.DB.Rites, *r)
	}
	err := repo.DB.saveToDisk()
	if err != nil {
		log.Printf("Could not save Rites to disk: %v", err)
	}
	return err
}

func (repo *RiteRepository) Get(title string) *domain.Rite {
	for i, r := range repo.DB.Rites {
		if title == r.Title {
			return &repo.DB.Rites[i]
		}
	}
	return nil
}

func (repo *RiteRepository) GetIds() []string {
	var ts []string
	for _, r := range repo.DB.Rites {
		ts = append(ts, r.Title)
	}
	return ts
}
