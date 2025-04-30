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

// TODO: membuat function Connect  

// TODO: membuat function CreateDB 

// TODO: membuat function DropDB 

// TODO: membuat function CreateTable 
