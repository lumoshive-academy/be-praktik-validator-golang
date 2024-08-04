package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	validator := validator.New()

	var username string = ""
	err := validator.Var(username, "required")

	if err != nil {
		fmt.Println("username validation failed:", err)
	} else {
		fmt.Println("username validation passed!")
	}

	// validation var with value
	ValidationVarWithValue()

	// multi tag validation
	MultiTagValidation()
}
