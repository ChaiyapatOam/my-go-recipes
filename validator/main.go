package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"email"`
}

func main() {
	user1 := User{Name: "John", Email: "123@sad"}
	validate := validator.New()
	err1 := validate.Struct(user1)
	if err1 != nil {
		fmt.Println("User Validate Error")
	}
	// fmt.Println("main validate")
}
