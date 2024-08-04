package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Definisikan struct dengan tag validasi
type Event struct {
	Name      string `validate:"required"`
	StartDate string `validate:"required,datetime=2006-01-02"`
	EndDate   string `validate:"required,datetime=2006-01-02,gtefield=StartDate"`
}

func ValidationCrossField() {
	// Inisialisasi validator
	validate := validator.New()

	// Contoh instance dari struct Event
	event := &Event{
		Name:      "Conference",
		StartDate: "2023-08-01",
		EndDate:   "2023-07-31", // EndDate lebih awal dari StartDate
	}

	// Validasi struct
	err := validate.Struct(event)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// Iterasi melalui errors dan tampilkan informasi detail
			for _, validationErr := range err.(validator.ValidationErrors) {
				fmt.Printf("Error: Field '%s' failed validation with tag '%s'.\n", validationErr.Field(), validationErr.Tag())
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
