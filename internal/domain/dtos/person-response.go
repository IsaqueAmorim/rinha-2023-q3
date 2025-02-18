package dtos

import "time"

type PersonResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate time.Time `json:"nascimento"`
	Stack     []string  `json:"stack"`
}
