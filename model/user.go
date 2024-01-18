package model

type User struct {
	UserId   string `json:"user_id" gorm:"primaryKey" gorm:"column:user_id" validate:"required"`
	Name     string `json:"name" gorm:"column:name" validate:"required"`
	Email    string `json:"email" gorm:"column:email" validate:"required,email"`
	Password string `json:"password" gorm:"column:password" validate:"required,min=8"`
	Phone    string `json:"phone" gorm:"column:phone"`
	UserType string `json:"user_type" gorm:"column:user_type"`
	Token    string `json:"token" gorm:"-"`
}
