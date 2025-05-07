package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// TODO: Membuat Struct Film
type Film struct {
	ID         int     `json:"id"`
	Judul      string  `json:"judul"`
	TahunRilis int     `json:"tahun_rilis"`
	Genre      string  `json:"genre"`
	Rating     float64 `json:"rating"`
}

// TODO: Implementasi dari slice
var films []Film

// TODO: Mengganti angka 81 pada PORT --> jadi 2 digit NPM terakhir
func main() {
	loadData()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respond(w, "Selamat datang di API Film! Gunakan /films untuk mengakses data film.", http.StatusOK)
	})
	http.HandleFunc("/films", handleFilms)
	http.HandleFunc("/films/", handleFilmByID)
	fmt.Println("Server berjalan di port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func loadData() { 															// Membaca file .json yang disimpan dalam function loadData
	file, err := os.ReadFile("films.json")
	if err == nil {
		json.Unmarshal(file, &films)
	}
}

func saveData() { 															// Menyimpan data ke .json
	data, _ := json.MarshalIndent(films, "", "  ")
	os.WriteFile("films.json", data, 0644)
}

// Method GET, POST
func handleFilms(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { 										// Data dapat di GET apabila sudah di POST
		respondJSON(w, films, http.StatusOK)
	} else if r.Method == http.MethodPost {
		var newFilm Film                                                 	// Data di POST agar dapat ditampilkan oleh server dengan GET
		if err := json.NewDecoder(r.Body).Decode(&newFilm); err == nil { 	// Membaca r.Body yang berisi data mentah di .json
			newFilm.ID = len(films) + 1    									// Memberikan ID baru otomatis ketika ada data baru
			films = append(films, newFilm) 									// Menyimpan ke dalam Slice
			saveData()                     									// Hasil akan disimpan pada function saveData
			respondJSON(w, newFilm, http.StatusCreated)
		} else {
			respond(w, "Format data tidak valid", http.StatusBadRequest)
		}
	} else {
		respond(w, "Metode tidak didukung", http.StatusMethodNotAllowed) 	// Output yang dihasilkan bila KALIAN running method "SELAIN" GET/POST
	}
}

// Method GET, PUT, DELETE
func handleFilmByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/films/"):]) 					// Mengambil ID dari URL path. Misalnya, /films/3 yang akan diambil "3"
	if err != nil {
		respond(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	for i, film := range films { 											// Melakukan looping pada slice films
		if film.ID == id { 													// Mencocokkan apakah film.ID sama dengan id di .json
			switch r.Method { 												// Mengecek HTTP Method Request
			case http.MethodGet: 											// If Method GET
				respondJSON(w, film, http.StatusOK)
			case http.MethodPut: 											// If Method PUT
				var updatedFilm Film
				if json.NewDecoder(r.Body).Decode(&updatedFilm) == nil {
					updatedFilm.ID = id
					films[i] = updatedFilm
					saveData()
					respondJSON(w, updatedFilm, http.StatusOK)
				} else {
					respond(w, "Format data tidak valid", http.StatusBadRequest)
				}
			case http.MethodDelete: 										// If Method DELETE
				films = append(films[:i], films[i+1:]...) 					// Menghapus film dari slice (men)
				saveData()
				respond(w, "data berhasil dihapus", http.StatusOK)
			default:
				respond(w, "Metode tidak didukung", http.StatusMethodNotAllowed)
			}
			return
		}
	}
	respond(w, "Film tidak ditemukan", http.StatusNotFound)
}

func respond(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
