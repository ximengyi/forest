package dao

import (
	"gorm.io/gorm"
	"time"
)


type Base struct {
	Id        uint           `json:"id" gorm:"primary_key" description:"自增主键"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"gorm:"index"`

}