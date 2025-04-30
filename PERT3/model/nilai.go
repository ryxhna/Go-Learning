package model

import (
	"database/sql"
	"fmt"
	"strings"
)

// TODO: membuat tabel nilai

// TODO: membuat struct nilai

func (m *Nilai) Structur() *Nilai {
	return &Nilai{}
}

func (m *Nilai) Fields() ([]string, []interface{}) {
	fields := []string{"id_nilai", "kd_mk", "npm", "uts", "uas", "total", "grade"}
	temp := []interface{}{&m.Id_nilai, &m.Kd_mk, &m.NPM, &m.UAS, &m.UTS, &m.Total, &m.Grade}
	return fields, temp
}

func (m *Nilai) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values(?,?,?,?,?,?,?)", "nilai")
	_, err := db.Exec(query, &m.Id_nilai, &m.Kd_mk, &m.NPM, &m.UAS, &m.UTS, &m.Total, &m.Grade)
	return err
}

func (m *Nilai) Update(db *sql.DB, data map[string]interface{}) error {

	var kolom = []string{}
	var args []interface{}

	i := 1

	for key, value := range data {
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
		i++
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "nilai", dataUpdate, "npm")
	args = append(args, m.NPM)
	_, err := db.Exec(query, args...)
	return err
}

func (m *Nilai) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "nilai", "npm")
	_, err := db.Exec(query, m.NPM)
	return err
}

// TODO: membuat function GetNilai dan GetAllNilai
