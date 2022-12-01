package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Title string
	Info  string
	Src   string
}
