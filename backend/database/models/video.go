package models

import (
	"gorm.io/gorm"
)

type Video struct {
	*gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
