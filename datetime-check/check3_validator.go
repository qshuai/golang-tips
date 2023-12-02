package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	checkValidator(&SearchUserV3Request{
		RegisterAt: "",
	})
}

type SearchUserV3Request struct {
	RegisterAt string `json:"register_at" validate:"omitempty,datetime=2006-01-02 15:04:05"`
}

func checkValidator(req *SearchUserV3Request) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
