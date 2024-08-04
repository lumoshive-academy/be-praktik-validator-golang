package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

// Definisikan struct dengan field map
type Inventory struct {
    Items map[string]int `validate:"dive,keys,required,endkeys,min=0"`
}

func main() {
    // Inisialisasi validator
    validate := validator.New()

    // Contoh instance dari struct Inventory dengan map Items
    inventory := &Inventory{
        Items: map[string]int{
            "item1": 10,
            "item2": -5,  // Nilai negatif, harus gagal validasi
            "item3": 15,
            "":      20,  // Key kosong, harus gagal validasi
        },
    }

    // Validasi struct
    err := validate.Struct(inventory)
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
