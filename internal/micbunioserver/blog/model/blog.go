package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string
	Content  string
	Views    int
	Slug     string
	Comments []Comment
}
