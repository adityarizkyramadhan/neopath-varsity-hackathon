package models

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name          string
	Vendor        string
	TotalQuestion int
	TimeEstimated time.Duration
	Description   string
	Instruction   string
}
