package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User adalah struktur data untuk pengguna
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`
}

// CreateUser menambahkan pengguna baru ke database
func CreateUser(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
    return err
}

// GetUserByUsername mendapatkan pengguna berdasarkan username
func GetUserByUsername(username string) (User, error) {
    var user User
    row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
    err := row.Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return user, err
    }
    return user, nil
}
