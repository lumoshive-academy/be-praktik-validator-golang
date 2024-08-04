package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Fungsi validasi khusus untuk memeriksa apakah field Password dan ConfirmPassword sama
func passwordMatch(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	confirmPasswordField, _, _, ok := fl.GetStructFieldOK2()
	if !ok {
		return false
	}
	confirmPassword := confirmPasswordField.String()
	return password == confirmPassword
}

// Definisikan struct dengan field yang menggunakan custom tag validasi
type User struct {
	Password        string `validate:"required"`
	ConfirmPassword string `validate:"required,password_match"` // Tag custom validation
}

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Daftarkan fungsi validasi khusus
	validate.RegisterValidation("password_match", passwordMatch)

	// Contoh instance dari struct User
	user := &User{
		Password:        "secret123",
		ConfirmPassword: "secret123", // Harus sama dengan Password untuk validasi berhasil
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
