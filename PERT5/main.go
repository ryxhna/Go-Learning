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

// TODO: Membuat function main
func main() {
	loadData()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respond(w, "Selamat datang di API Film! Gunakan /films untuk mengakses data film.", http.StatusOK)
	})
	http.HandleFunc("/films", handleFilms)
	http.HandleFunc("/films/", handleFilmByID)
	fmt.Println("Server berjalan di port 8007")
	log.Fatal(http.ListenAndServe(":8007", nil))
}

func loadData() {
	file, err := os.ReadFile("films.json")
	if err == nil {
		json.Unmarshal(file, &films)
	}
}

func saveData() {
	data, _ := json.MarshalIndent(films, "", "  ")
	os.WriteFile("films.json", data, 0644)
}

// Method GET, POST ---------------------------------------------------------------------------------
// Data dapat di GET apabila sudah di POST
func handleFilms(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		respondJSON(w, films, http.StatusOK)
	} else if r.Method == http.MethodPost {
		var newFilm Film
		if err := json.NewDecoder(r.Body).Decode(&newFilm); err == nil {
			newFilm.ID = len(films) + 1
			films = append(films, newFilm)
			saveData()
			respondJSON(w, newFilm, http.StatusCreated)
		} else {
			respond(w, "Format data tidak valid", http.StatusBadRequest)
		}
	} else {
		respond(w, "Metode tidak didukung", http.StatusMethodNotAllowed)
	}
}

// Method GET, PUT, DELETE --------------------------------------------------------------------------
// Mengambil ID dari URL path. Misalnya, /films/3 yang akan diambil "3"
// Melakukan looping pada slice films kemudian mencocokkan apakah film.ID sama dengan id di .json

func handleFilmByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/films/"):])
	if err != nil {
		respond(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	for i, film := range films {
		if film.ID == id {
			switch r.Method {
			case http.MethodGet:
				respondJSON(w, film, http.StatusOK)
			case http.MethodPut:
				var updatedFilm Film
				if json.NewDecoder(r.Body).Decode(&updatedFilm) == nil {
					updatedFilm.ID = id
					films[i] = updatedFilm
					saveData()
					respondJSON(w, updatedFilm, http.StatusOK)
				} else {
					respond(w, "Format data tidak valid", http.StatusBadRequest)
				}
			case http.MethodDelete:
				films = append(films[:i], films[i+1:]...)
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

// TODO: Membuat Respons untuk JSON
func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
