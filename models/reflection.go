package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Question string `gorm:"not null"`
	Section  int
	Role     string
	Page     int
}

type Answer struct {
	gorm.Model
	StudentID  uint `gorm:"not null"`
	QuestionID uint `gorm:"not null;foreignKey:QuestionID"`
	Answer     int  `gorm:"not null"`
}

type AnswerInput struct {
	QuestionID int
	Answer     int
}

type Evaluation struct {
	Average    float64 `gorm:"not null"`
	QuestionID uint    `gorm:"not null;foreignKey:QuestionID"`
	StudentID  uint    `gorm:"not null;foreignKey:StudentID"`
}
