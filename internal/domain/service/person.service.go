package service

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/entity"
	postgr "github.com/IsaqueAmorim/rinha-2023/internal/infra/database"
)

func CreatePerson(p *entity.Person) (string, error) {

	if err := p.Validate(); err != nil {
		return "", err
	}

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

// func GetPersonById(id string) (*entity.Person, error) {
// 	p, err := r.FindByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return p, nil
// }

// func CountPersons() (int, error) {
// 	count, err := r.Count()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }
