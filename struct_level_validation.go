package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

// Struct dengan beberapa field yang perlu divalidasi secara keseluruhan
type User struct {
    Password        string `validate:"required"`
    ConfirmPassword string `validate:"required"`
    Age             int    `validate:"gte=0,lte=130"`
}

// Fungsi validasi untuk seluruh struct User
func userStructLevelValidation(sl validator.StructLevel) {
    user := sl.Current().Interface().(User)

    // Memastikan bahwa Password dan ConfirmPassword harus sama
    if user.Password != user.ConfirmPassword {
        sl.ReportError(user.ConfirmPassword, "ConfirmPassword", "ConfirmPassword", "passwordmatch", "")
    }

    // Contoh validasi lainnya: memastikan usia harus lebih dari 18 tahun jika password diisi
    if user.Password != "" && user.Age < 18 {
        sl.ReportError(user.Age, "Age", "Age", "agegte18", "")
    }
}

func main() {
    // Inisialisasi validator
    validate := validator.New()

    // Daftarkan fungsi validasi khusus untuk struct User
    validate.RegisterStructValidation(userStructLevelValidation, User)

    // Contoh instance dari struct User
    user := &User{
        Password:        "secret123",
        ConfirmPassword: "secret123",
        Age:             17, // Ini akan gagal karena usia kurang dari 18 tahun
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
