package cmd

import "github.com/tdekeyser/rite/core/domain"

type Module struct {
	db domain.RiteRepository
}

func NewModule(db domain.RiteRepository) *Module {
	return &Module{db: db}
}

func (m *Module) SaveRite(title string, body string) error {
	r := &domain.Rite{Title: title, Body: []byte(body)}
	return m.db.Save(r)
}

func (m *Module) GetRite(title string) *domain.Rite {
	r := m.db.Get(title)
	if r == nil {
		return &domain.Rite{Title: title}
	}
	return r
}

func (m *Module) GetAllRiteTitles() []string {
	return m.db.GetIds()
}
