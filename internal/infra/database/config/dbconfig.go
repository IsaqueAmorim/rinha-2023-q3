package config

import "fmt"

type DBConfig struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DbName  string
	MaxPool int
	MinPool int
}

func GetDefaultConfig() DBConfig {
	return DBConfig{
		Host:    "db",
		Port:    5432,
		User:    "postgres",
		Pass:    "admin",
		MaxPool: 100,
		MinPool: 5,
		DbName:  "rinha",
	}
}

func GetDefaultConnectionString() string {
	dbconfig := GetDefaultConfig()

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Pass, dbconfig.DbName)

	return connectionString
}
