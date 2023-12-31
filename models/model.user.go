package model

import (
	"time"

	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityUsers struct {
	ID        string         `gorm:"primaryKey;"`
	Username  string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);unique;not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Photos    []EntityPhotos `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = helper.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
