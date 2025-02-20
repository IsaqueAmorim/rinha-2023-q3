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
	Nickname  string    `json:"apelido" validate:"required, max=32"`
	Name      string    `json:"nome" validate:"required, max=100"`
	Birthdate time.Time `json:"nascimento" validate:"required,datetime=2006-01-02"`
	Stack     []string  `json:"stack" validate:"dive,max=32"`
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
	dateLayout := "2006-01-02"
	if _, err := time.Parse(dateLayout, p.Birthdate.Format(dateLayout)); err != nil {
		return errors.New("Invalid date format")
	}

	return nil
}
