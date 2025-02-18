package repository

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/entity"
)

type PersonRepository interface {
	Insert(p *entity.Person) (string, error)
	// FindAll() (*entity.Person, error)
	// FindByID(id string) (*entity.Person, error)
	// Count() (int, error)
	// FindByText(text string) ([]entity.Person, error)
}
