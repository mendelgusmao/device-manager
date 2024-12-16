package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
)

type devicesMap map[string]*models.Device

type DeviceRepositoryMock struct {
	devices devicesMap
}

func NewDeviceRepositoryMock() *DeviceRepositoryMock {
	return &DeviceRepositoryMock{
		devices: make(devicesMap),
	}
}

func (r *DeviceRepositoryMock) Insert(_ context.Context, device models.Device) (*models.Device, error) {
	storedDevice := &models.Device{}
	copier.Copy(&storedDevice, &device)
	storedDevice.ID = uuid.NewString()

	r.devices[storedDevice.ID] = storedDevice

	return storedDevice, nil
}

func (r *DeviceRepositoryMock) FetchOne(_ context.Context, query *models.DeviceQuery) (*models.Device, error) {
	device, ok := r.devices[*query.ID]

	if !ok {
		return nil, fmt.Errorf("fetchOne: `%s` not found", *query.ID)
	}

	return device, nil
}

func (r *DeviceRepositoryMock) FetchMany(context.Context, *models.DeviceQuery) ([]models.Device, error) {
	devices := make([]models.Device, 0)

	for _, device := range r.devices {
		devices = append(devices, *device)
	}

	return devices, nil
}
func (r *DeviceRepositoryMock) Update(_ context.Context, device models.Device) (*models.Device, error) {
	fetchedDevice, ok := r.devices[device.ID]

	if !ok {
		return nil, fmt.Errorf("update: `%s` not found", device.ID)
	}

	copier.Copy(&fetchedDevice, &device)

	return fetchedDevice, nil
}

func (r *DeviceRepositoryMock) Delete(_ context.Context, query models.DeviceQuery) error {
	_, ok := r.devices[*query.ID]

	if !ok {
		return fmt.Errorf("delete: `%s` not found", *query.ID)
	}

	delete(r.devices, *query.ID)

	return nil
}

func (r *DeviceRepositoryMock) Setup() error {
	return nil
}
