package v1

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	"github.com/mendelgusmao/device-manager/internal/tests/mocks/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	s := service.NewDeviceServiceMock()
	h := NewDeviceHandlers(s)
	ctx := context.Background()

	createDevice := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	response, err := h.createDevice(ctx, &struct{ Body dtos.CreateDevice }{
		Body: createDevice,
	})

	device := response.Body

	assert.NoError(t, err)
	assert.NotEmpty(t, device.ID)
	assert.Equal(t, createDevice.BrandName, device.BrandName)
	assert.Equal(t, createDevice.DeviceName, device.DeviceName)
	assert.NotZero(t, device.CreatedAt.Unix())
}

func TestGetDevice(t *testing.T) {
	s := service.NewDeviceServiceMock()
	h := NewDeviceHandlers(s)
	ctx := context.Background()
	id := uuid.NewString()

	response, err := h.getDevice(ctx, &deviceIdentifier{
		Identifier: id,
	})

	assert.NoError(t, err)

	device := response.Body

	assert.Equal(t, id, device.ID)
	assert.NotEmpty(t, device.BrandName)
	assert.NotEmpty(t, device.DeviceName)
	assert.NotZero(t, device.CreatedAt.Unix())
}

func TestListDevices(t *testing.T) {
	s := service.NewDeviceServiceMock()
	h := NewDeviceHandlers(s)
	ctx := context.Background()

	createDevice := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	response, err := h.listDevices(ctx, nil)

	assert.NoError(t, err)

	devices := response.Body

	assert.Len(t, devices, 1)

	device := devices[0]
	assert.NotEmpty(t, device.ID)
	assert.Equal(t, createDevice.BrandName, device.BrandName)
	assert.Equal(t, createDevice.DeviceName, device.DeviceName)
	assert.NotZero(t, device.CreatedAt.Unix())
}

func TestSearchDevices(t *testing.T) {
	s := service.NewDeviceServiceMock()
	h := NewDeviceHandlers(s)
	ctx := context.Background()

	createDevice := dtos.CreateDevice{
		BrandName:  "Test Brand",
		DeviceName: "Test Device",
	}

	response, err := h.searchDevices(ctx, &deviceSearch{})

	assert.NoError(t, err)

	devices := response.Body

	assert.Len(t, devices, 1)

	device := devices[0]
	assert.NotEmpty(t, device.ID)
	assert.Equal(t, createDevice.BrandName, device.BrandName)
	assert.Equal(t, createDevice.DeviceName, device.DeviceName)
	assert.NotZero(t, device.CreatedAt.Unix())
}

func TestUpdateDevices(t *testing.T) {
	s := service.NewDeviceServiceMock()
	h := NewDeviceHandlers(s)
	ctx := context.Background()

	id := uuid.NewString()
	brandName := "Another Test Brand"
	deviceName := "Another Test Device"

	updateDevice := dtos.UpdateDevice{
		BrandName:  &brandName,
		DeviceName: &deviceName,
	}

	response, err := h.updateDevice(ctx, &deviceUpdateRequest{
		Identifier: id,
		Body:       updateDevice,
	})

	assert.NoError(t, err)

	device := response.Body

	assert.Equal(t, id, device.ID)
	assert.Equal(t, *updateDevice.BrandName, device.BrandName)
	assert.Equal(t, *updateDevice.DeviceName, device.DeviceName)
	assert.NotZero(t, device.CreatedAt.Unix())
}

func TestDeleteDevices(t *testing.T) {
	s := service.NewDeviceServiceMock()
	h := NewDeviceHandlers(s)
	ctx := context.Background()

	id := uuid.NewString()

	_, err := h.deleteDevice(ctx, &deviceIdentifier{
		Identifier: id,
	})

	assert.NoError(t, err)
}
