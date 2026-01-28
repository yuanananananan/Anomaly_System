package models

import "time"

type BaseModel struct {
	ID        uint64    `gorm:"primaryKey:autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updatedAt"`
}
