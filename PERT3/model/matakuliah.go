package model

import (
	"database/sql"
	"fmt"
)

// ---------------------- BAGIAN INI TIDAK PERLU DI UBAH --------------------------- //
// --------------------------------------------------------------------------------- //
var TabelMatkul string = `
	CREATE TABLE matkul(
		kd_mk VARCHAR(10) PRIMARY KEY,
		matakuliah VARCHAR(20)
	)	
`

type Matkul struct {
	Kd_mk      string `json:"Kd_mk"`
	Matakuliah string `json:"Matakuliah"`
}

// --------------------------------------------------------------------------------- //
// TODO: membuat isi dari tabel --> struktur, fields, insert

// ---------------------- BAGIAN INI TIDAK PERLU DI UBAH --------------------------- //
// --------------------------------------------------------------------------------- //
func GetMatkul(db *sql.DB, id string) (*Matkul, error) {

	m := &Matkul{}
	each := m.Structur()
	_, dst := each.Fields()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", "matkul", "kd_mk")
	err := db.QueryRow(query, id).Scan(dst...)
	if err != nil {
		return nil, err
	}
	return each, nil
}

func GetAllMatkul(db *sql.DB) ([]*Matkul, error) {
	m := &Matkul{}
	query := fmt.Sprintf("SELECT * FROM %s", "matkul")
	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	var result []*Matkul

	for data.Next() {
		each := m.Structur()
		_, dst := each.Fields()
		err := data.Scan(dst...)

		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}
