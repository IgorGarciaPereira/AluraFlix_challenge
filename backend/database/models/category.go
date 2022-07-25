package models

import (
  "gorm.io/gorm"
	_ "gopkg.in/validator.v2"
)

type Category struct{
  gorm.Model

  Title string  `json:"title" validate:"nonzero"`
  Color string  `json:"color" validate:"nonzero"`
}
