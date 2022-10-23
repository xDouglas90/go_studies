package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, nonzero, regexp=^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[\-]?[0-9]{2}$, unique"`
	RG   string `json:"rg" validate:"min=9, max=10, nonzero, regexp=^[0-9]+$"`
}

func Validate(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
