package main

import (
    "fmt"
    "regexp"

    "github.com/go-playground/validator/v10"
)

// Fungsi validasi khusus untuk memeriksa format username
func usernameValidation(fl validator.FieldLevel) bool {
    username := fl.Field().String()
    // Contoh: username harus mengandung huruf kecil dan angka saja
    match, _ := regexp.MatchString("^[a-z0-9]+$", username)
    return match
}

// Definisikan struct dengan field yang menggunakan custom tag validasi
type User struct {
    Username string `validate:"username"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"gte=0,lte=130"`
}

func main() {
    // Inisialisasi validator
    validate := validator.New()

    // Daftarkan fungsi validasi khusus
    validate.RegisterValidation("username", usernameValidation)

    // Contoh instance dari struct User
    user := &User{
        Username: "john_doe", // Username ini harus gagal validasi
        Email:    "johndoe@example.com",
        Age:      25,
    }

    // Validasi struct
    err := validate.Struct(user)
    if err != nil {
        if _, ok := err.(validator.ValidationErrors); ok {
            // Iterasi melalui errors dan tampilkan informasi detail
            for _, validationErr := range err.(validator.ValidationErrors) {
                fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Namespace(), validationErr.Tag())
                fmt.Printf("  Value: '%v'\n", validationErr.Value())
                fmt.Printf("  Condition: '%s'\n", validationErr.Param())
            }
        } else {
            // Error lainnya
            fmt.Println("Validation failed:", err)
        }
    } else {
        fmt.Println("Validation passed!")
    }
}
