package models

import "gorm.io/gorm"

type MetaLearningPath struct {
	gorm.Model
	Name             string `gorm:"unique;not null"`
	Role             string
	Deskripsi        string
	Level            string
	DataLearningPath []DataLearningPath `gorm:"foreignKey:MetaID"`
}

type MetaLearningPathInput struct {
	Name      string
	Role      string
	Deskripsi string
	Level     string
}

type DataLearningPath struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Url       string `gorm:"not null"`
	Thumbnail string
	IsPaid    string           `gorm:"not null"`
	Source    string           `gorm:"not null"`
	MetaID    uint             `gorm:"index;not null"`
	MetaLP    MetaLearningPath `gorm:"foreignKey:MetaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DataLearningPathInput struct {
	Title     string
	Url       string
	Thumbnail string
	IsPaid    string
	Source    string
	MetaID    uint
}

type StudentProgress struct {
	gorm.Model
	IsDone             bool             `gorm:"default:false;not null"`
	StudentID          uint             `gorm:"not null"`
	MetaLearningPathID uint             `gorm:"not null"`
	Student            Student          `gorm:"foreignKey:StudentID"`
	MetaLearningPath   MetaLearningPath `gorm:"foreignKey:MetaLearningPathID"`
}
