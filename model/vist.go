package model

import (
	"github.com/lib/pq"
	"time"
)

type Visit struct {
	VisitId   string    `json:"visit_id" gorm:"primaryKey" gorm:"column:visit_id" validate:"required"`
	UserId    string    `json:"user_id" gorm:"column:user_id"`
	Doctor    string    `json:"doctor" gorm:"column:doctor"`
	Hospital  string    `json:"hospital" gorm:"column:hospital"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
	Type      string    `json:"type" gorm:"column:type"`
	Comment   string    `json:"comment" gorm:"column:comment"`
	/*	DoctorId  string         `json:"doctor_id" gorm:"column:doctor_id"`
	 */Drugs  pq.StringArray `json:"drugs" gorm:"type:text[]" gorm:"column:drugs"`
	Reports   pq.StringArray `json:"reports" gorm:"type:text[]" gorm:"column:reports"`
	ScanFiles pq.StringArray `json:"scan_files" gorm:"type:text[]" gorm:"column:scan_files"`
}
