package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	Nama      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
