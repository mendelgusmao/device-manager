package services

import (
	"context"
	"testing"

	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	mocks "github.com/mendelgusmao/device-manager/internal/tests/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	r := mocks.NewDeviceRepositoryMock()
	s := NewDeviceService(r)
	ctx := context.Background()

	device := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	newDevice, err := s.CreateDevice(ctx, device)

	assert.NoError(t, err)
	assert.NotEmpty(t, newDevice.ID)
	assert.Equal(t, device.BrandName, newDevice.BrandName)
	assert.Equal(t, device.DeviceName, newDevice.DeviceName)
}

func TestGetDevice(t *testing.T) {
	r := mocks.NewDeviceRepositoryMock()
	s := NewDeviceService(r)
	ctx := context.Background()

	device := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	newDevice, err := s.CreateDevice(ctx, device)

	assert.NoError(t, err)

	fetchedDevice, err := s.GetDevice(ctx, newDevice.ID)

	assert.Equal(t, newDevice.ID, fetchedDevice.ID)
	assert.Equal(t, newDevice.BrandName, fetchedDevice.BrandName)
	assert.Equal(t, newDevice.DeviceName, fetchedDevice.DeviceName)
}

func TestGetDevices(t *testing.T) {
	r := mocks.NewDeviceRepositoryMock()
	s := NewDeviceService(r)
	ctx := context.Background()

	device := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	newDevice, err := s.CreateDevice(ctx, device)

	assert.NoError(t, err)

	fetchedDevices, err := s.GetDevices(ctx, nil)

	assert.NotEmpty(t, fetchedDevices)

	fetchedDevice := fetchedDevices[0]
	assert.Equal(t, newDevice.ID, fetchedDevice.ID)
	assert.Equal(t, newDevice.BrandName, fetchedDevice.BrandName)
	assert.Equal(t, newDevice.DeviceName, fetchedDevice.DeviceName)
}

func TestUpdateDevice(t *testing.T) {
	r := mocks.NewDeviceRepositoryMock()
	s := NewDeviceService(r)
	ctx := context.Background()

	device := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	newDevice, err := s.CreateDevice(ctx, device)

	assert.NoError(t, err)
	assert.NotEmpty(t, newDevice.ID)

	newBrandName := "New Test Brand"
	newDeviceName := "New Device Name"

	updateDevice := dtos.UpdateDevice{
		BrandName:  &newBrandName,
		DeviceName: &newDeviceName,
	}

	updatedDevice, err := s.UpdateDevice(ctx, newDevice.ID, updateDevice)
	assert.NoError(t, err)

	assert.Equal(t, newDevice.ID, updatedDevice.ID)
	assert.Equal(t, *updateDevice.BrandName, updatedDevice.BrandName)
	assert.Equal(t, *updateDevice.DeviceName, updatedDevice.DeviceName)
}

func TestDeleteDevice(t *testing.T) {
	r := mocks.NewDeviceRepositoryMock()
	s := NewDeviceService(r)
	ctx := context.Background()

	device := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	newDevice, err := s.CreateDevice(ctx, device)

	assert.NoError(t, err)
	assert.NotEmpty(t, newDevice.ID)

	err = s.DeleteDevice(ctx, newDevice.ID)
	assert.NoError(t, err)

	_, err = s.GetDevice(ctx, newDevice.ID)

	assert.Error(t, err)
}
