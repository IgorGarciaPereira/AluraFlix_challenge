package models

import (
	"gorm.io/gorm"
	_ "gopkg.in/validator.v2"
)

type Video struct {
	gorm.Model

	Title       string `json:"title" validate:"nonzero"`
	Description string `json:"description" validate:"nonzero"`
	Url         string `json:"url" validate:"nonzero"`
}
