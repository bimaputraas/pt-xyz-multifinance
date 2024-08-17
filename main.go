package main

import (
	"xyz-multifinance/pkg"
)

type User struct {
	Email string `validate:"required,email"`
}

func main() {
	pkg.ValidateStruct(User{Email: "ss"})

}
