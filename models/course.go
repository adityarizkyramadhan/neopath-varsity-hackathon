package models

import "gorm.io/gorm"

type MetaLearningPath struct {
	gorm.Model
	Name             string `gorm:"unique;not null"`
	IsDone           bool
	Role             string
	DataLearningPath []DataLearningPath `gorm:"foreignKey:MetaID"`
}

type DataLearningPath struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Url       string `gorm:"not null"`
	Thumbnail string
	IsPaid    string           `gorm:"not null"`
	Source    string           `gorm:"not null"`
	MetaID    uint             `gorm:"index;not null"`
	MetaLP    MetaLearningPath `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
