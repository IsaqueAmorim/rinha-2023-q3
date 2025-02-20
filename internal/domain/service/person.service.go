package service

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/entity"
	postgr "github.com/IsaqueAmorim/rinha-2023/internal/infra/database"
)

func CreatePerson(p *entity.Person) (string, error) {

	id, err := postgr.Insert(p)
	if err != nil {
		return "", err
	}

	return id, nil
}

func CheckIfNicknameExists(nickname string) (bool, error) {
	exists, err := postgr.CheckIfNicknameExists(nickname)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func GetPersonById(id string) (*entity.Person, error) {
	p, err := postgr.FindByID(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func CountPersons() (int, error) {
	count, err := postgr.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func FindByText(text string) ([]*entity.Person, error) {
	persons, err := postgr.FindByText(text)
	if err != nil {
		return nil, err
	}

	return persons, nil
}
