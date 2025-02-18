package postgr

import (
	"database/sql"

	"github.com/IsaqueAmorim/rinha-2023/internal/infra/database/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connectDB()
	createTable()
}

func GetConnection() *sql.DB {
	return db
}

func connectDB() {
	con, err := sql.Open("postgres", config.GetDefaultConnectionString())

	if err != nil {
		panic(err)
	}

	if err = con.Ping(); err != nil {
		panic(err)
	}
	db = con
}

func createTable() {
	_, err := db.Exec(`
    CREATE TABLE 
      IF NOT EXISTS pessoas 
      (
        id uuid PRIMARY KEY NOT NULL,
        nickname varchar(32) UNIQUE NOT NULL,
        name varchar(100) NOT NULL,
        birthdate date NOT NULL,
        stack varchar(255)[] NULL
      )`)

	if err != nil {
		panic(err)
	}
}
