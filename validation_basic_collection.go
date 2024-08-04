package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan struct dengan field koleksi
type Order struct {
	Items []Item `validate:"required,dive"`
}

type Item struct {
	Name  string `validate:"required"`
	Price int    `validate:"gte=0"`
}

func ValidationMap() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh instance dari struct Order dengan koleksi Items
	order := &Order{
		Items: []Item{
			{Name: "Item1", Price: 100},
			{Name: "", Price: 200},      // Nama item ini kosong, harus gagal validasi
			{Name: "Item3", Price: -10}, // Harga item ini negatif, harus gagal validasi
		},
	}

	// Validasi struct
	err := validate.Struct(order)
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
