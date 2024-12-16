package devices

import (
	"context"

	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
)

type Repository interface {
	Insert(context.Context, models.Device) (*models.Device, error)
	FetchOne(context.Context, *models.DeviceQuery) (*models.Device, error)
	FetchMany(context.Context, *models.DeviceQuery) ([]models.Device, error)
	Update(context.Context, models.Device) (*models.Device, error)
	Delete(context.Context, models.DeviceQuery) error
	Setup() error
}
