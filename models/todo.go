package models

import (
	"time"
)

type Todo struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
