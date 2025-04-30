package test

import (
	"PERT3/model"
	"testing"
)

func TestMahasiswa(t *testing.T) {
	var dataInsertMhs = []model.Mahasiswa{
		model.Mahasiswa{
			NPM:   "00000000",
			Nama:  "Widjoyo",
			Kelas: "4KA00",
		},
		model.Mahasiswa{
			NPM:   "11111111",
			Nama:  "Admin1",
			Kelas: "4KA19",
		},
		model.Mahasiswa{
			NPM:   "22222222",
			Nama:  "Admin2",
			Kelas: "4KA99",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert mahasiswa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMhs {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update mahasiswa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"nama": "Abdi Teh Ayeuna",
		}

		data := dataInsertMhs[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Get mahasiswa", func(t *testing.T) {
		_, err := model.GetMahasiswa(db, "11111111")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing Get mahasiswa", func(t *testing.T) {
		_, err := model.GetAllMahasiswa(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing delete mahasiswa", func(t *testing.T) {
		data := dataInsertMhs[0]
		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
	defer db.Close()
}
