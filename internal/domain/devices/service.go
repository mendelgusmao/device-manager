package devices

import (
	"context"

	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
)

type Service interface {
	CreateDevice(context.Context, dtos.CreateDevice) (*dtos.Device, error)
	GetDevice(context.Context, string) (*dtos.Device, error)
	GetDevices(context.Context, *models.DeviceQuery) ([]dtos.Device, error)
	UpdateDevice(context.Context, string, dtos.UpdateDevice) (*dtos.Device, error)
	DeleteDevice(context.Context, string) error
}
