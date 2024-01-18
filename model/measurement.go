package model

import (
	"time"
)

type Measurement struct {
	MeasurementId string    `json:"measurement_id" gorm:"primaryKey" gorm:"column:user_id" validate:"required"`
	UserId        string    `json:"user_id" gorm:"column:user_id"`
	Type          string    `json:"type" gorm:"column:type"`
	XValue        string    `json:"x_value" gorm:"column:x_value"`
	YValue        string    `json:"y_value" gorm:"column:y_value"`
	Name          string    `json:"name" gorm:"column:name"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	TakenBy       string    `json:"taken_by" gorm:"column:taken_by"`
}
