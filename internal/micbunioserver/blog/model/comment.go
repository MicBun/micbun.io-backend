package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	BlogID  uint
	Content string
	Name    string
	Blog    Blog
}
