package model

import (
	"database/sql"
	"fmt"
	"strings"
)

// TODO: membuat tabel matakuliah

// TODO: membuat struct matkul

func (m *Matkul) Structur() *Matkul {
	return &Matkul{}
}

func (m *Matkul) Fields() ([]string, []interface{}) {
	fields := []string{"kd_mk", "matakuliah"}
	temp := []interface{}{&m.Kd_mk, &m.Matakuliah}
	return fields, temp
}

func (m *Matkul) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v VALUES(?,?)", "matkul")
	_, err := db.Exec(query, m.Kd_mk, m.Matakuliah)
	return err
}

func (m *Matkul) Update(db *sql.DB, data map[string]interface{}) error {

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
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "matkul", dataUpdate, "kd_mk")
	args = append(args, m.Kd_mk)
	_, err := db.Exec(query, args...)
	return err
}

func (m *Matkul) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "matkul", "kd_mk")
	_, err := db.Exec(query, m.Kd_mk)
	return err
}

// TODO: membuat function GetMatkul dan GetAllMatkul
