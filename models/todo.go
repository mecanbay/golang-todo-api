package models

import (
	"time"
)

type Todo struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null;default:null" json:"title"`
	Desc        string    `json:"desc" gorm:"default:null"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
