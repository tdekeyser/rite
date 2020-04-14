package filestorage

import "github.com/tdekeyser/rite/core/domain"

type TagRepository struct {
	DB *dataStore
}

func NewTagRepository(db *dataStore) *TagRepository {
	return &TagRepository{DB: db}
}

func (repo *TagRepository) Create(t *domain.Tag) error {
	repo.DB.Tags[*t] = true
	return nil
}

func (repo *TagRepository) GetAll() []domain.Tag {
	tags := make([]domain.Tag, len(repo.DB.Tags))
	for t := range repo.DB.Tags {
		tags = append(tags, t)
	}
	return tags
}
