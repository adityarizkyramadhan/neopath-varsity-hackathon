package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Test struct {
		gorm.Model
		Name                string
		Vendor              string
		TotalQuestion       int
		TimeEstimated       time.Duration
		Description         string
		Instruction         string
		StudentSurveyScores []StudentSurveyScore
	}

	StudentSurveyScore struct {
		gorm.Model
		StudentID             uint
		SchoolID              uint
		TestID                uint
		ScoreProblemSolving   int
		ScoreCriticalThinking int
		ScoreBusinessAcumen   int
		ScoreDetailOriented   int
		ScoreCommunication    int
		ScoreDataAnalysis     int
	}

	SurveyQuestion struct {
		gorm.Model
		Q1              string // Do you enjoy working in a team?
		Q2              string // Are you comfortable with public speaking?
		Q3              string // Do you prefer working with numbers over words?
		Q4              string // Are you comfortable with ambiguity and uncertainty?
		Q5              string // Do you enjoy problem-solving and critical thinking?
		SurveyResponses []SurveyResponse
	}

	SurveyResponse struct {
		gorm.Model
		StudentID        uint
		SurveyQuestionID uint
		Q1Answer         string // Answer to "Do you enjoy working in a team?"
		Q2Answer         string // Answer to "Are you comfortable with public speaking?"
		Q3Answer         string // Answer to "Do you prefer working with numbers over words?"
		Q4Answer         string // Answer to "Are you comfortable with ambiguity and uncertainty?"
		Q5Answer         string // Answer to "Do you enjoy problem-solving and critical thinking?"
	}
)
