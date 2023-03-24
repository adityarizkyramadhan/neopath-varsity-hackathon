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
	SchoolID        uint   `gorm:"ForeignKey:SchoolID"`
	School          School `gorm:"AssociationForeignKey:ID"`
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
