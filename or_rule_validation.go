package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

// Definisikan struct dengan field yang menggunakan tag validasi dengan rule or
type User struct {
    Contact string `validate:"required,email|numeric"` // Contact bisa berupa email atau angka
}

func main() {
    // Inisialisasi validator
    validate := validator.New()

    // Contoh instance dari struct User
    user := &User{
        Contact: "1234567890", // Ini adalah nilai numerik yang valid
        // Contact: "user@example.com", // Ini adalah email yang valid
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
