package models

import "time"

type Article struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(200)" binding:"required,min=20"`
	Content     string    `json:"content" gorm:"type:text" binding:"required,min=200"`
	Category    string    `json:"category" gorm:"type:varchar(100)" binding:"required,min=3"`
	Status      string    `json:"status" gorm:"type:varchar(100)" binding:"required,oneof=publish draft thrash"`
	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}