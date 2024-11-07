
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

type User struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name,omitempty"`
    Address  string `json:"address,omitempty"`
}

// Fungsi untuk membaca data pengguna dari file JSON
func readUsers() ([]User, error) {
    var users []User
    if _, err := os.Stat("users.json"); err == nil {
        data, err := ioutil.ReadFile("users.json")
        if err != nil {
            return nil, err
        }
        err = json.Unmarshal(data, &users)
        if err != nil {
            return nil, err
        }
    }
    return users, nil
}

// Fungsi untuk menyimpan data pengguna ke file JSON
func saveUsers(users []User) error {
    data, err := json.MarshalIndent(users, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("users.json", data, 0644)
}

// Endpoint untuk registrasi
func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Hanya POST yang diizinkan", http.StatusMethodNotAllowed)
        return
    }

    var newUser User
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, "Gagal membaca input", http.StatusBadRequest)
        return
    }

    users, err := readUsers()
    if err != nil {
        http.Error(w, "Gagal membaca data pengguna", http.StatusInternalServerError)
        return
    }

    // Periksa apakah pengguna sudah terdaftar
    for _, user := range users {
        if user.Email == newUser.Email {
            http.Error(w, "Pengguna sudah terdaftar", http.StatusBadRequest)
            return
        }
    }

    users = append(users, newUser)
    if err := saveUsers(users); err != nil {
        http.Error(w, "Gagal menyimpan data pengguna", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Registrasi berhasil"))
}

// Endpoint untuk login
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Hanya POST yang diizinkan", http.StatusMethodNotAllowed)
        return
    }

    var loginUser User
    if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
        http.Error(w, "Gagal membaca input", http.StatusBadRequest)
        return
    }

    users, err := readUsers()
    if err != nil {
        http.Error(w, "Gagal membaca data pengguna", http.StatusInternalServerError)
        return
    }

    // Verifikasi email dan password
    for _, user := range users {
        if user.Email == loginUser.Email && user.Password == loginUser.Password {
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Login berhasil"))
            return
        }
    }

    http.Error(w, "Email atau password salah", http.StatusUnauthorized)
}

func main() {
    // Menyajikan file statis dari folder "static"
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Menyajikan halaman registrasi (index.html)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    // Endpoint untuk registrasi
    http.HandleFunc("/register", registerHandler)

    // Endpoint untuk login
    http.HandleFunc("/login", loginHandler)

    fmt.Println("Server berjalan di http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
