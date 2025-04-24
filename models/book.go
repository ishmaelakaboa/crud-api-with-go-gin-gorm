package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Year   int    `json:"year"`
}

func (Book) TableName() string{
	return "books"
}