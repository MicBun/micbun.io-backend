package model

import "gorm.io/gorm"

type Guestbook struct {
	gorm.Model
	Name    string
	Content string
	HostURL string
}
