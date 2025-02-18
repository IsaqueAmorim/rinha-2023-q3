package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/IsaqueAmorim/rinha-2023/internal/domain/dtos"
	"github.com/google/uuid"
)

type Person struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate time.Time `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

func NewPerson(p *dtos.PersonRequest) *Person {

	person := &Person{
		ID:        uuid.NewString(),
		Nickname:  p.Nickname,
		Name:      p.Name,
		Birthdate: p.Birthdate,
		Stack:     p.Stack,
	}

	person.Validate()

	return person
}

func (p *Person) Validate() error {

	if strings.TrimSpace(p.Name) == "" {
		return errors.New("Name is required")
	}
	if strings.TrimSpace(p.Nickname) == "" {
		return errors.New("Nickname is required")
	}

	return nil
}
