package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan struct dengan tag validasi
type User struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=0,lte=130"`
	Password string `validate:"required,min=8"`
	Address  string `validate:"required"`
	Phone    string `validate:"required,e164"` // e164 adalah format internasional untuk nomor telepon
}

func ValidationStruct() {

	// Inisialisasi validator
	validate := validator.New()

	// Contoh instance dari struct User
	user := &User{
		Name:     "Lumoshive Academy",
		Email:    "lumoshive@example.com",
		Age:      28,
		Password: "password123",
		Address:  "Grand Garden",
		Phone:    "+628709090930",
	}

	// // Validasi struct
	// err := validate.Struct(user)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("Validation passed!")
	// }

	// Validasi struct
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: Field %s failed validation with tag %s\n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("Validation passed!")
	}

}
