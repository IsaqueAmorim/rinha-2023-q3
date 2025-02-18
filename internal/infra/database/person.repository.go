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
