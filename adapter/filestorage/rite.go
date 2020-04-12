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
	} else {
		repo.DB.rites = append(repo.DB.rites, *r)
	}
	err := repo.DB.saveToDisk()
	if err != nil {
		log.Printf("Could not save rites to disk: %v", err)
	}
	return err
}

func (repo *RiteRepository) Get(title string) *domain.Rite {
	for i, r := range repo.DB.rites {
		if title == r.Title {
			return &repo.DB.rites[i]
		}
	}
	return nil
}

func (repo *RiteRepository) GetIds() []string {
	var ts []string
	for _, r := range repo.DB.rites {
		ts = append(ts, r.Title)
	}
	return ts
}
