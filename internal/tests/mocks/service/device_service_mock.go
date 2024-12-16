package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
)

type DeviceServiceMock struct{}

func NewDeviceServiceMock() *DeviceServiceMock {
	return &DeviceServiceMock{}
}

func (s *DeviceServiceMock) CreateDevice(ctx context.Context, device dtos.CreateDevice) (*dtos.Device, error) {
	deviceDTO := dtos.Device{
		ID: uuid.NewString(),
	}

	copier.Copy(&deviceDTO, &device)

	return &deviceDTO, nil
}

func (s *DeviceServiceMock) GetDevice(ctx context.Context, identifier string) (*dtos.Device, error) {
	return &dtos.Device{
		ID:         identifier,
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}, nil
}

func (s *DeviceServiceMock) GetDevices(ctx context.Context, filter *models.DeviceQuery) ([]dtos.Device, error) {
	return []dtos.Device{
		{
			ID:         uuid.NewString(),
			BrandName:  "Test Brand",
			DeviceName: "Test Device",
		},
	}, nil
}

func (s *DeviceServiceMock) UpdateDevice(ctx context.Context, identifier string, device dtos.UpdateDevice) (*dtos.Device, error) {
	deviceDTO := &dtos.Device{
		ID: identifier,
	}

	copier.Copy(&deviceDTO, device)

	return deviceDTO, nil
}

func (s *DeviceServiceMock) DeleteDevice(ctx context.Context, identifier string) error {
	return nil
}
