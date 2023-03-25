package models

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Name       string
	Occupation string
	Rating     int
}

type MentorRegister struct {
	Name       string
	Occupation string
	Rating     int
}
