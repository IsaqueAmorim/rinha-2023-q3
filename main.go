package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Pessoa struct {
	Apelido    string
	Nome       string
	Nascimento string
	Stack      []string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "pessoas"
)

var db *sql.DB

func main() {
	connectDB()
	criarTabela()
}

func connectDB() {
	// Conecta ao banco de dados
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	con, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	db = con
	fmt.Println("Conectado ao banco de dados")
}

func criarTabela() {
	// Cria a tabela
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS pessoas (id PRIMARY KEY,apelido varchar(255) UNIQUE, nome varchar(255), nascimento date, stack varchar(255)[])")

	if err != nil {
		panic(err)
	}

	fmt.Println("Tabela criada")
}

func inserirPessoa(pessoa *Pessoa) {
	if err := validarPessoa(pessoa); err != nil {
		panic(err)
	}

	_, err := db.Exec("INSERT INTO pessoas (apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4)",
		pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pessoa.Stack)

	if err != nil {
		panic(err)
	}
}

func validarPessoa(pessoa *Pessoa) error {
	if strings.TrimSpace(pessoa.Apelido) == "" {
		return fmt.Errorf("apelido não pode ser vazio")
	}
	if strings.TrimSpace(pessoa.Nome) == "" {
		return fmt.Errorf("nome não pode ser vazio")
	}

	return nil
}
