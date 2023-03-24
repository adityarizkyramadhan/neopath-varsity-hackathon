package models

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Name       string
	Occupation string
	Rating     string
}

type MentorRegister struct {
	Name       string
	Occupation string
	Rating     string
}
