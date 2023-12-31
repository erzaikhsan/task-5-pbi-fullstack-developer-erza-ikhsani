package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityPhotos struct {
	ID        string `gorm:"primaryKey;"`
	Title     string `gorm:"type:varchar(255);not null"`
	Caption   string `gorm:"type:varchar(255);not null"`
	PhotoUrl  string `gorm:"type:varchar(255);not null"`
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityPhotos) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityPhotos) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
