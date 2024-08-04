package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan nested struct
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Zip    string `validate:"required"`
}

// Definisikan struct utama yang memiliki nested struct
type User struct {
	Name    string  `validate:"required"`
	Email   string  `validate:"required,email"`
	Age     int     `validate:"gte=0,lte=130"`
	Address Address `validate:"required,dive"`
}

func ValidationNestedStruct() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh instance dari struct User dengan nested struct Address
	user := &User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Age:   30,
		Address: Address{
			Street: "123 Main St",
			City:   "",
			Zip:    "12345",
		},
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
