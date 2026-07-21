// ! cmd/app/main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"baji-golang-in-memory-database/internal/cache"
)

func MainRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hellowCoy := "Welcome"
	json.NewEncoder(w).Encode(hellowCoy)

}

func main() {
	// 1. Inisialisasi Cache Generics kustom Anda untuk tipe data string
	// Petugas kebersihan (Janitor) disetel patroli setiap 5 menit sesuai instruksi
	kustomCache := cache.New[string](5 * time.Minute)

	// 2. Mengisi data uji coba dengan TTL 10 menit
	kustomCache.Set("misi_hari_6", "Arsitektur HTTP & Cache Berhasil!", 10*time.Minute)

	// 3. Setup HTTP Multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", MainRoute)
	// 4. Registrasi Endpoint API
	mux.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Mengambil data dari internal cache
		if data, ditemukan := kustomCache.Get("misi_hari_6"); ditemukan {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"status": "success", "data": "%s"}`, data)
			return
		}

		// Response jika data kedaluwarsa atau tidak ada
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"status": "error", "message": "data tidak ditemukan atau sudah basi"}`)
	})

	// 5. Jalankan Web Server
	fmt.Println("Server berjalan di http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}
