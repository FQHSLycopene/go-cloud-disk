package models

import (
	"gorm.io/gorm"
	"time"
)

type Name struct {
	ID        int64          `gorm:"column:id" json:"id"`
	Identity  string         `gorm:"column:identity" json:"identity"`
	Name      string         `gorm:"column:name" json:"name"`
	Password  string         `gorm:"column:password" json:"password"`
	Email     string         `gorm:"column:email" json:"email"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (t *Name) TableName() string {
	return "name"
}
