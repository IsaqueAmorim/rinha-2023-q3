package postgr

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func Insert(p *entity.Person) (string, error) {
	con, err := GetConnection().Begin()

	if err != nil {
		return "", err
	}

	query := `
	INSERT INTO pessoas 
	(
		id,
		name,
		nickname,
		birthdate,
		stack
	) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`

	stmt, err := con.Prepare(query)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	var id string
	err = stmt.QueryRow(uuid.NewString(), p.Name, p.Nickname, p.Birthdate, pq.Array(p.Stack)).Scan(&id)
	if err != nil {
		con.Rollback()
		return "", err
	}

	con.Commit()
	return id, nil

}

func CheckIfNicknameExists(nickname string) (bool, error) {
	con, err := GetConnection().Begin()

	if err != nil {
		return false, err
	}

	query := `SELECT COUNT(1) FROM pessoas WHERE nickname = $1;`

	stmt, err := con.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var count int
	err = stmt.QueryRow(nickname).Scan(&count)
	if err != nil {
		con.Rollback()
		return false, err
	}

	con.Commit()
	return count > 0, nil
}

func FindByID(id string) (*entity.Person, error) {
	con, err := GetConnection().Begin()

	if err != nil {
		return nil, err
	}

	query := `SELECT id, name, nickname, birthdate, stack FROM pessoas WHERE id = $1;`

	stmt, err := con.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p entity.Person
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Nickname, &p.Birthdate, pq.Array(&p.Stack))
	if err != nil {
		con.Rollback()
		return nil, err
	}

	con.Commit()
	return &p, nil
}

func Count() (int, error) {
	con, err := GetConnection().Begin()

	if err != nil {
		return 0, err
	}

	query := `SELECT COUNT(1) FROM pessoas;`

	stmt, err := con.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var count int
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		con.Rollback()
		return 0, err
	}

	con.Commit()
	return count, nil
}

func FindByText(text string) ([]*entity.Person, error) {
	con, err := GetConnection().Begin()

	if err != nil {
		return nil, err
	}

	query := `SELECT id, name, nickname, birthdate, stack FROM pessoas WHERE name ILIKE $1 OR nickname ILIKE $1;`

	stmt, err := con.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(text)
	if err != nil {
		con.Rollback()
		return nil, err
	}

	defer rows.Close()

	var persons = make([]*entity.Person, 0)
	for rows.Next() {
		var p entity.Person
		err = rows.Scan(&p.ID, &p.Name, &p.Nickname, &p.Birthdate, pq.Array(&p.Stack))
		if err != nil {
			con.Rollback()
			return nil, err
		}
		persons = append(persons, &p)
	}

	con.Commit()
	return persons, nil
}
