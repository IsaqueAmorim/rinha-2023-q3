package entity

import (
	"fmt"
	"strings"
)

type Pessoa struct {
  ID int
	Apelido    string
	Nome       string
	Nascimento string
	Stack      []string
}

func (p *Pessoa) validate() error {
  if strings.TrimSpace(p.Nome) == ""{
    return fmt.Errorf("Nome não pode ser vazio!")
  }
  if strings.TrimSpace(p.Apelido) == "" {
    return fmt.Errorf("Apelido não pode ser Vazio")
  }

  return nil
}
