package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan struct dengan field map
type Product struct {
	Attributes map[string]string `validate:"dive,keys,required,endkeys,required"`
}

func main() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh instance dari struct Product dengan map Attributes
	product := &Product{
		Attributes: map[string]string{
			"Color":  "Red",
			"Size":   "L",
			"Weight": "",                // Value ini kosong, harus gagal validasi
			"":       "ValueWithoutKey", // Key ini kosong, harus gagal validasi
		},
	}

	// Validasi struct
	err := validate.Struct(product)
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
