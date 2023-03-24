package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Email           string
	Password        string `json:"-"`
	Name            string
	Gender          string
	Age             int
	SelfDescription string
	SchoolID        uint   `gorm:"ForeignKey:SchoolID"`
	School          School `gorm:"AssociationForeignKey:ID"`
}

type StudentDTO struct {
	gorm.Model
	Email           string
	Password        string `json:"-"`
	Name            string
	Gender          string
	Age             int
	SelfDescription string
}

type StudentUpdate struct {
	Name            string
	Gender          string
	Age             int
	SelfDescription string
}

type StudentLogin struct {
	Email    string
	Password string
}

type StudentRegister struct {
	Email    string
	Password string
	SchoolID uint
}
