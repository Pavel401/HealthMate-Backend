package model

import "time"

type Public_Page struct {
	PageId   string    `json:"page_id" gorm:"primaryKey" gorm:"column:file_id" validate:"required"`
	UserId   string    `json:"user_id" gorm:"column:user_id"`
	Code     string    `json:"code" gorm:"column:code"`
	ValidTil time.Time `json:"valid_til" gorm:"column:valid_till"`
}
