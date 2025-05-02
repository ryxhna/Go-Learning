# Pertemuan 4

## Setup SQL

1. Makesure your directory is
   `C:\go\src\PERT4>`
2. Installation Go Modules
   `go mod init PERT4`
3. Installation MySQL Driver
   `go get github.com/go-sql-driver/mysql`
4. Check your location now bt running this command in terminal
   `go env GOPATH`
5. 
    `go mod tidy`

## Changes for your code

1. main.go: membuat function main, pastikan port sudah sesuai
2. model/database.go: pastikan port sudah sesuai
3. test/database_test.go: pastikan port sudah sesuai
4. test/mahasiswa_test: isi dengan data kalian
5. handler/handleMatkul.go: copy function dari file handlerMahasiswa.go, sesuaikan dengan perintah TODO didalamnya
6. handler/handleNilai.go: copy function dari file handlerMahasiswa.go, sesuaikan dengan perintah TODO didalamnya
7. handler/handle.go: pastikan port sudah sesuai
8. handler/handle.go: membuat switch case untuk Mahasiswa, Nilai, Matkul
