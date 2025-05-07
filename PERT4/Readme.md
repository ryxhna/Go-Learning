# Pertemuan 4

## Changes for your code (general)
1. main.go: membuat function main, pastikan port sudah sesuai
2. model/database.go: pastikan port sudah sesuai
3. test/database_test.go: pastikan port sudah sesuai
4. test/mahasiswa_test: isi dengan data kalian

### handleMatkul.go (spesific)
TODO:
- Membuat function HandlerMatkulGet
- Membuat function HandlerMatkulPost
- Membuat function HandlerMatkulDelete
- Membuat function HandlerMatkulPut
- Mengubah npm --> kd_mk
- Mengubah NPM --> Kd_mk (perhatikan kapital) == line 52, 81
- Ganti kata "Mahasiswa" --> "Matkul"

### handleNilai.go (spesific)
TODO:
- Membuat function HandlerNilaiGet
- Membuat function HandlerNilaiPost
- Membuat function HandlerNilaiDelete
- Membuat function HandlerNilaiPut
- Ganti kata "Mahasiswa" --> "Nilai"

## handler.go (spesific)
- TODO: Sesuaikan Nama DB dengan --> db_PERT4_GO
- TODO: Switch Case untuk Nilai == perhatikan kapital atau tidaknya
- TODO: Switch Case untuk Matkul == perhatikan kapital atau tidaknya
- Jangan lupa setelah case matkul copy & paste lagi untul defaultnya seperti apa
   default:
      w.Write([]byte("request tidak ditemukan"))
   }

## Setup SQL
1. Makesure your directory is
   `C:\go\src\PERT4>`
2. Installation Go Modules
   `go mod init PERT4`
3. Installation MySQL Driver
   `go get github.com/go-sql-driver/mysql`
4. Check your location now bt running this command in terminal
   `go env GOPATH`
5. Search packet & refresh automatically
   `go mod tidy`


## Running your code
1. Open XAMPP
   - Turn on MySQL & Apache 
   - Open Shell
   - mysql -u root
   - show databases;

2. Open XAMPP 
   - cd test
   - go test -run TestDatabase
   - go test -run TestMahasiswa

3. Open Terminal (VSCODE)
   - use db_PERT4_GO 
   - show tables;
   - select * from mahasiswa;
   
4. Open Browser & Terminal (VSCODE)
   - Browser = http://localhost/phpmyadmin == Cek IP, misalnya http://127.0.0.1/
   - VSCODE  = cd .. 
   - VSCODE  = go run main.go
   - Browser = http://127.0.0.1:8007/api/mahasiswa


# Screenshots
1. SS hasil terminal setelah setup  SQL
2. SS folder kalian (break down folder handler)
3. SS hasil terminal VSCODE == TestDatabase -== TestMahasiswa
4. SS hasil == show tables; == select * from mahasiswa;
5. SS hasil open http://127.0.0.1:8007/api/mahasiswa

