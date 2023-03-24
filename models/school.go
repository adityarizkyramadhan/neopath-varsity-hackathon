package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Name     string
	Students []Student `gorm:"ForeignKey:SchoolID"`
}
