package models

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Name       string
	Occupation string
	Rating     string
}
