package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Mengatur handler untuk rute utama "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Menyajikan file statis (misalnya gambar, CSS, dll.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Menjalankan server pada port 8080
	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
