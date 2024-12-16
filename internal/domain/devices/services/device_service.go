package services

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/mendelgusmao/device-manager/internal/domain/devices"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
)

type DeviceService struct {
	repository devices.Repository
}

func NewDeviceService(repository devices.Repository) *DeviceService {
	return &DeviceService{
		repository: repository,
	}
}

func (s *DeviceService) CreateDevice(ctx context.Context, device dtos.CreateDevice) (*dtos.Device, error) {
	newDevice, err := s.repository.Insert(ctx, models.Device{
		BrandName:  device.BrandName,
		DeviceName: device.DeviceName,
	})

	if err != nil {
		return nil, err
	}

	deviceDTO := &dtos.Device{}
	copier.Copy(deviceDTO, newDevice)

	return deviceDTO, nil
}

func (s *DeviceService) GetDevice(ctx context.Context, identifier string) (*dtos.Device, error) {
	device, err := s.repository.FetchOne(ctx, &models.DeviceQuery{
		ID: &identifier,
	})

	if err != nil {
		return nil, err
	}

	deviceDTO := &dtos.Device{}
	copier.Copy(deviceDTO, device)

	return deviceDTO, nil
}

func (s *DeviceService) GetDevices(ctx context.Context, filter *models.DeviceQuery) ([]dtos.Device, error) {
	devices, err := s.repository.FetchMany(ctx, filter)

	if err != nil {
		return nil, err
	}

	fetchedDevices := make([]dtos.Device, len(devices))
	copier.Copy(&fetchedDevices, devices)

	return fetchedDevices, nil
}

func (s *DeviceService) UpdateDevice(ctx context.Context, identifier string, device dtos.UpdateDevice) (*dtos.Device, error) {
	deviceModel := models.Device{
		ID: identifier,
	}

	copier.Copy(&deviceModel, device)
	updatedDevice, err := s.repository.Update(ctx, deviceModel)

	if err != nil {
		return nil, err
	}

	deviceDTO := dtos.Device{}
	copier.Copy(&deviceDTO, updatedDevice)

	return &deviceDTO, nil
}

func (s *DeviceService) DeleteDevice(ctx context.Context, identifier string) error {
	return s.repository.Delete(ctx, models.DeviceQuery{ID: &identifier})
}
