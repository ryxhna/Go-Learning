### Import package yang diperlukan

- encoding/json : konversi data ke/dari format JSON
- fmt, log : output dan logging
- net/http : membuat HTTP server
- os : file handling
- strconv : konversi string ke int

## Penjelasan

- handleFilms = http/localhost:8080/films --> GET, POST
- handleFilmByID = http/localhost:8080/films/{id} --> GET, PUT, DELETE
- loadData = Membaca file .json yang disimpan dalam function loadData
- saveData = Menyimpan data ke .json

## Setup SQL

1. Makesure your directory is
   `C:\go\src\PERT5>`
2. Installation Go Modules
   `go mod init PERT5`
3. Installation MySQL Driver
   `go get github.com/go-sql-driver/mysql`
4. Check your location now bt running this command in terminal
   `go env GOPATH`
5. Search packet & refresh automatically
   `go mod tidy`

## Your task

1. main.go = Membuat Struct Film
2. main.go = Implementasi dari slice
3. main.go = Membuat function main
4. main.go = Membuat Respons untuk JSON

## How to running?

1. xampp = Click start for "Apache"
2. open terminal = go mod init PERT5
3. open terminal = go run main.go
4. open chrome = http://localhost:8007 --> screenshoots (samakan dengan port kaliann)
5. open chrome = http://localhost:8007/films --> screenshoots
6. open chrome = http://localhost:8007/films/ 1 --> screenshoots
