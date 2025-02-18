package dtos

import "time"

type PersonRequest struct {
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate time.Time `json:"nascimento"`
	Stack     []string  `json:"stack"`
}
