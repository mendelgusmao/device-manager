package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	ID         string `gorm:"primaryKey"`
	BrandName  string `gorm:"column:brand_name"`
	DeviceName string `gorm:"column:device_name"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (device *Device) BeforeCreate(tx *gorm.DB) error {
	if device.ID == "" {
		device.ID = uuid.NewString()
	}

	return nil
}
