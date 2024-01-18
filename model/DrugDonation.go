package model

import (
	"github.com/lib/pq"
)

type Drug_Donation struct {
	DonationId string         `json:"donation_id" gorm:"primaryKey" gorm:"column:donation_id" validate:"required"`
	UserId     string         `json:"user_id" gorm:"column:user_id"`
	Status     string         `json:"status" gorm:"column:status"`
	DrugName   string         `json:"drug_name" gorm:"column:drug_name"`
	Pictures   pq.StringArray `json:"pictures" gorm:"type:text[]" gorm:"column:drugs"`
	ExpireDate string         `json:"expire_date" gorm:"column:expire_date"`
	Category   string         `json:"category" gorm:"column:category"`
}
