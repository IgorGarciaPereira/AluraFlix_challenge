package models

import (
  "gorm.io/gorm"
	_ "gopkg.in/validator.v2"
)

type Category struct{
  gorm.Model

  ID    int     `json:"id"`
  Title string  `json:"title"`
  Color string  `json:"color"`
}
