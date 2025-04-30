package model

import (
	"database/sql"
	"fmt"
	"strings"
)

// TODO: membuat tabel mahasiswa

// TODO: membuat struct mahasiswa

func (m *Mahasiswa) Fields() ([]string, []interface{}) {
	fields := []string{"npm", "nama", "kelas"}
	temp := []interface{}{&m.NPM, &m.Nama, &m.Kelas}
	return fields, temp
}

func (m *Mahasiswa) Structur() *Mahasiswa {
	return &Mahasiswa{}
}

func (m *Mahasiswa) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values(?,?,?)", "mahasiswa")
	_, err := db.Exec(query, m.NPM, m.Nama, m.Kelas)
	return err
}

func (m *Mahasiswa) Update(db *sql.DB, data map[string]interface{}) error {

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
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "mahasiswa", dataUpdate, "NPM")
	args = append(args, m.NPM)
	_, err := db.Exec(query, args...)
	return err
}

func (m *Mahasiswa) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "mahasiswa", "NPM")
	_, err := db.Exec(query, m.NPM)
	return err
}

// TODO: membuat function GetMahasiswa dan GetAllMahasiswa
