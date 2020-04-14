package domain

type Tag string

type TagRepository interface {
	Create(*Tag) error
	GetAll() []Tag
}
