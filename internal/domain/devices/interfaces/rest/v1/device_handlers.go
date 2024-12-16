package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/mendelgusmao/device-manager/internal/domain/devices"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/errors"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
)

type DeviceHandlers struct {
	service devices.Service
}

func NewDeviceHandlers(service devices.Service) *DeviceHandlers {
	return &DeviceHandlers{
		service: service,
	}
}

func (h DeviceHandlers) createDevice(ctx context.Context, input *struct{ Body dtos.CreateDevice }) (*deviceResponse, error) {
	device, err := h.service.CreateDevice(ctx, input.Body)

	if err != nil {
		return nil, err
	}

	return &deviceResponse{
		Body: device,
	}, nil
}

func (h DeviceHandlers) getDevice(ctx context.Context, input *deviceIdentifier) (*deviceResponse, error) {
	device, err := h.service.GetDevice(ctx, input.Identifier)

	if err == errors.ErrorRecordNotFound {
		return nil, huma.Error404NotFound(
			fmt.Sprintf("device `%s` not found", input.Identifier),
			err,
		)
	}

	if err != nil {
		return nil, err
	}

	return &deviceResponse{
		Body: device,
	}, err
}

func (h DeviceHandlers) listDevices(ctx context.Context, input *struct{}) (*devicesResponse, error) {
	devices, err := h.service.GetDevices(ctx, nil)

	if err != nil {
		return nil, err
	}

	return &devicesResponse{
		Body: devices,
	}, nil
}

func (h DeviceHandlers) searchDevices(ctx context.Context, input *deviceSearch) (*devicesResponse, error) {
	query := models.DeviceQuery{}

	if input.BrandName != "" {
		query.BrandName = &input.BrandName
	}

	if input.DeviceName != "" {
		query.DeviceName = &input.DeviceName
	}

	devices, err := h.service.GetDevices(ctx, &query)

	if err != nil {
		return nil, err
	}

	return &devicesResponse{
		Body: devices,
	}, nil
}

func (h DeviceHandlers) updateDevice(ctx context.Context, input *deviceUpdateRequest) (*deviceResponse, error) {
	device, err := h.service.UpdateDevice(ctx, input.Identifier, input.Body)

	if err == errors.ErrorRecordNotFound {
		return nil, huma.Error404NotFound(
			fmt.Sprintf("device `%s` not found", input.Identifier),
			err,
		)
	}

	if err != nil {
		return nil, err
	}

	return &deviceResponse{
		Body: device,
	}, nil
}

func (h DeviceHandlers) deleteDevice(ctx context.Context, input *deviceIdentifier) (*struct{}, error) {
	err := h.service.DeleteDevice(ctx, input.Identifier)

	if err == errors.ErrorRecordNotFound {
		return nil, huma.Error404NotFound(
			fmt.Sprintf("device `%s` not found", input.Identifier),
			err,
		)
	}

	return nil, err
}

func (h DeviceHandlers) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "add-device",
		Method:        http.MethodPost,
		Path:          "/api/v1/devices",
		Summary:       "Add a device",
		DefaultStatus: http.StatusCreated,
	}, h.createDevice)

	huma.Register(api, huma.Operation{
		OperationID: "search-devices",
		Method:      http.MethodGet,
		Path:        "/api/v1/devices/search",
		Summary:     "Search devices",
	}, h.searchDevices)

	huma.Register(api, huma.Operation{
		OperationID: "get-device",
		Method:      http.MethodGet,
		Path:        "/api/v1/devices/{identifier}",
		Summary:     "Get a device by identifier",
	}, h.getDevice)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/api/v1/devices",
		Summary: "List all devices",
	}, h.listDevices)

	huma.Register(api, huma.Operation{
		OperationID: "update-device",
		Method:      http.MethodPatch,
		Path:        "/api/v1/devices/{identifier}",
		Summary:     "Update a device",
	}, h.updateDevice)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-device",
		Method:        http.MethodDelete,
		Path:          "/api/v1/devices/{identifier}",
		Summary:       "Delete a device",
		DefaultStatus: http.StatusNoContent,
	}, h.deleteDevice)
}
