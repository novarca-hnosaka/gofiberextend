package gofiber_extend_test

import (
	"testing"

	gofiber_extend "github.com/novarca-hnosaka/gofiber_exntend"
)

func TestSimpleValidation(t *testing.T) {
	ex := gofiber_extend.New(gofiber_extend.IFiberExConfig{})
	err := ex.SimpleValidation("abc", "test", "match=^[a-z]+$")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

type StructTest struct {
	Name  string `validate:"required,match=^[a-z]+$"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,min=18,max=30"`
	Sex   int    `validate:"required"`
}

func TestValidation(t *testing.T) {
	ex := gofiber_extend.New(gofiber_extend.IFiberExConfig{})
	src := StructTest{
		Name:  "qwerty",
		Email: "hoge@hoge.com",
		Age:   20,
		Sex:   1,
	}
	err := ex.Validation(&src)
	if err != nil {
		t.Errorf("%+v", err)
	}
}
