package test

import (
	"PERT4/model"
	"testing"
)

func TestMahasiswa(t *testing.T) {

	var dataInsertMhs = []model.Mahasiswa{
		model.Mahasiswa{
			NPM:   "12345678",
			Nama:  "Agus",
			Kelas: "3KA20",
		},
		model.Mahasiswa{
			NPM:   "19283746",
			Nama:  "Andre",
			Kelas: "4KA20",
		},
		// TODO: Isi data kalian 
		model.Mahasiswa{
			NPM:   "10121507",
			Nama:  "Ghina Desrizkymalia Zahirah",
			Kelas: "4KA09",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert  mahasiswa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMhs {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update  mahasiswa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			// UBAH NAMA BERIKUT
			"nama": "Agus",
		}

		data := dataInsertMhs[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Get mahasiswa", func(t *testing.T) {		
		_, err := model.GetMahasiswa(db, "12345678")
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
