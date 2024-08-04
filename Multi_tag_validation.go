package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MultiTagValidation() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh validasi string dengan multi tag
	username := "lumoshive"
	err := validate.Var(username, "required,alphanum")
	if err != nil {
		fmt.Println("Username validation failed:", err)
	} else {
		fmt.Println("Username validation passed!")
	}

	// Contoh validasi angka dengan multi tag
	age := 25
	err = validate.Var(age, "gte=18,lte=65")
	if err != nil {
		fmt.Println("Age validation failed:", err)
	} else {
		fmt.Println("Age validation passed!")
	}
}
