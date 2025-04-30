package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Table interface {
	Name() string
	Field() ([]string, []interface{})
}

func Connect(username string, password string, host string, database string) (*sql.DB, error) {
	// TODO: sesuaikan dengan port mysql kalian
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, database)
	db, err := sql.Open("mysql", conn)
	return db, err
}

// TODO: membuat function CreateDB, CreateTable, DropDB
