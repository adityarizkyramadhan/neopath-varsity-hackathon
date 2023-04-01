package models

import (
	"gorm.io/gorm"
)

type (
	ApptitudeTest struct {
		gorm.Model
		Question string
		Section  int
		AnswerA  string
		AnswerB  string
		AnswerC  string
		AnswerD  string
	}

	ApptitudeAnswer struct {
		gorm.Model
		StudentID       uint
		ApptitudeTestID uint
		Answer          string
	}
)
