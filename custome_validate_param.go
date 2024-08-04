package main

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// Fungsi validasi khusus untuk memeriksa panjang minimal string
func minLength(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := fl.Param() // Mendapatkan parameter dari tag validasi
	minLength, err := strconv.Atoi(param)
	if err != nil {
		return false // Jika parameter tidak valid, validasi gagal
	}
	return len(value) >= minLength
}

// Definisikan struct dengan field yang menggunakan custom tag validasi dengan parameter
type User struct {
	Username string `validate:"minlength=12"` // Username harus memiliki panjang minimal 5 karakter
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=0,lte=130"`
}

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Daftarkan fungsi validasi khusus dengan parameter
	validate.RegisterValidation("minlength", minLength)

	// Contoh instance dari struct User
	user := &User{
		Username: "lumoshive", // Username ini harus gagal validasi karena panjangnya kurang dari 5 karakter
		Email:    "lumoshive@example.com",
		Age:      20,
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
