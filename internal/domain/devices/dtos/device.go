package dtos

import "time"

type Device struct {
	ID         string    `json:"id"`
	BrandName  string    `json:"brandName"`
	DeviceName string    `json:"deviceName"`
	CreatedAt  time.Time `json:"createdAt"`
}
