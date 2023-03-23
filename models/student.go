package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name            string
	Email           string
	Password        string
	Gender          string
	Age             string
	SelfDescription string
}
