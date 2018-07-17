package utils

import (
	"testing"
	"fmt"
)

type User struct {
	Name string `valid:"NotEmpty str-max-len=10" json:"name"`
	Age  int `valid:"int-max=100" json:"age"`
}

func TestValidate(t *testing.T) {
	user := &User{}
	err := StructValidate(user)
	fmt.Println(err)
	user = &User{"hello world",23}
	err = StructValidate(user)
	fmt.Println(err)
	user = &User{"hello",101}
	err = StructValidate(user)
	fmt.Println(err)
}
