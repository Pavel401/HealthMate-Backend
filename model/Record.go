package model

import "time"

type Aws_File struct {
	FileId    string    `json:"file_id" gorm:"primaryKey" gorm:"column:file_id" validate:"required"`
	UserId    string    `json:"user_id" gorm:"column:user_id"`
	AwsLink   string    `json:"aws_link" gorm:"column:aws_link"`
	Extension string    `json:"extension" gorm:"column:extension"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
