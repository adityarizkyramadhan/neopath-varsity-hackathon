package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Question string `gorm:"not null"`
	Section  int
	Role     string
	Page     int
}

type QuestionInput struct {
	Question string
	Section  int
	Role     string
	Page     int
}

type Answer struct {
	gorm.Model
	StudentID  uint `gorm:"not null"`
	QuestionID uint `gorm:"not null;foreignKey:QuestionID"`
	SectionID  uint
	Answer     int `gorm:"not null"`
}

type AnswerInput struct {
	QuestionID int
	Answer     int
	SectionID  uint
}

type Evaluation struct {
	Empathetic      float64
	Analytical      float64
	Adaptive        float64
	Collaborative   float64
	DesignSensitive float64
	AvgAll          float64
}
