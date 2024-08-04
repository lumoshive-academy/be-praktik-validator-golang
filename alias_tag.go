package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan struct dengan field yang menggunakan alias tag
type User struct {
	Username string `validate:"username_alias"`
	Email    string `validate:"email_alias"`
	Age      int    `validate:"age_alias"`
}

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Daftarkan alias tag
	validate.RegisterAlias("username_alias", "required,alphanum,min=3,max=32")
	validate.RegisterAlias("email_alias", "required,email")
	validate.RegisterAlias("age_alias", "gte=0,lte=130")

	// Contoh instance dari struct User
	user := &User{
		Username: "JohnDoe",
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
