package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"
	v1 "github.com/mendelgusmao/device-manager/internal/domain/devices/interfaces/rest/v1"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/repository"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/services"
	"github.com/mendelgusmao/device-manager/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func setupAPI(t *testing.T) humatest.TestAPI {
	db, err := database.NewSQLiteDatabase("file::memory:")

	assert.NoError(t, err)

	r := repository.NewDeviceRepository(db)
	err = r.Setup()
	assert.NoError(t, err)

	s := services.NewDeviceService(r)
	h := v1.NewDeviceHandlers(s)

	_, api := humatest.New(t)

	h.RegisterRoutes(api)

	return api
}

func TestPostDevice(t *testing.T) {
	api := setupAPI(t)
	ctx := context.Background()

	requestBody := map[string]any{
		"brandName":  "Test Brand",
		"deviceName": "Test Device",
	}

	response := api.PostCtx(ctx, "/api/v1/devices", requestBody)

	device := dtos.Device{}
	err := json.NewDecoder(response.Body).Decode(&device)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, requestBody["brandName"], device.BrandName)
	assert.Contains(t, requestBody["deviceName"], device.DeviceName)
	assert.NotZero(t, device.CreatedAt.Unix())
}

func TestGetDevice(t *testing.T) {
	api := setupAPI(t)
	ctx := context.Background()

	requestBody := map[string]any{
		"brandName":  "Test Brand",
		"deviceName": "Test Device",
	}

	response := api.PostCtx(ctx, "/api/v1/devices", requestBody)

	createdDevice := dtos.Device{}
	err := json.NewDecoder(response.Body).Decode(&createdDevice)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)

	response = api.GetCtx(ctx, fmt.Sprintf("/api/v1/devices/%s", createdDevice.ID))

	device := dtos.Device{}
	err = json.NewDecoder(response.Body).Decode(&device)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, createdDevice, device)
}

func TestGetDevices(t *testing.T) {
	api := setupAPI(t)
	ctx := context.Background()

	requestBody := map[string]any{
		"brandName":  "Test Brand",
		"deviceName": "Test Device",
	}

	response := api.PostCtx(ctx, "/api/v1/devices", requestBody)

	createdDevice := dtos.Device{}
	err := json.NewDecoder(response.Body).Decode(&createdDevice)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)

	response = api.GetCtx(ctx, "/api/v1/devices", createdDevice.ID)

	devices := []dtos.Device{}
	err = json.NewDecoder(response.Body).Decode(&devices)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, []dtos.Device{createdDevice}, devices)
}

func TestPatchDevice(t *testing.T) {
	api := setupAPI(t)
	ctx := context.Background()

	requestBody := map[string]any{
		"brandName":  "Test Brand",
		"deviceName": "Test Device",
	}

	response := api.PostCtx(ctx, "/api/v1/devices", requestBody)

	createdDevice := dtos.Device{}
	err := json.NewDecoder(response.Body).Decode(&createdDevice)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)

	updateRequestBody := map[string]any{
		"brandName": "Other Test Brand",
	}

	response = api.PatchCtx(
		ctx,
		fmt.Sprintf("/api/v1/devices/%s", createdDevice.ID),
		updateRequestBody,
	)

	device := dtos.Device{}
	err = json.NewDecoder(response.Body).Decode(&device)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotEqual(t, createdDevice, device)
}

func TestDeleteDevice(t *testing.T) {
	api := setupAPI(t)
	ctx := context.Background()

	requestBody := map[string]any{
		"brandName":  "Test Brand",
		"deviceName": "Test Device",
	}

	response := api.PostCtx(ctx, "/api/v1/devices", requestBody)

	createdDevice := dtos.Device{}
	err := json.NewDecoder(response.Body).Decode(&createdDevice)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)

	response = api.DeleteCtx(
		ctx,
		fmt.Sprintf("/api/v1/devices/%s", createdDevice.ID),
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestSearchDevice(t *testing.T) {
	api := setupAPI(t)
	ctx := context.Background()

	postDevice := func(body map[string]any) *dtos.Device {
		response := api.PostCtx(ctx, "/api/v1/devices", body)

		createdDevice := dtos.Device{}
		err := json.NewDecoder(response.Body).Decode(&createdDevice)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, response.Code)

		return &createdDevice
	}

	createdDevice := postDevice(map[string]any{
		"brandName":  "TestBrand",
		"deviceName": "TestDevice",
	})

	postDevice(map[string]any{
		"brandName":  "Another Test Brand",
		"deviceName": "Another Test Device",
	})

	response := api.GetCtx(ctx, "/api/v1/devices/search?brandName=TestBrand")

	devices := []dtos.Device{}
	err := json.NewDecoder(response.Body).Decode(&devices)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, []dtos.Device{*createdDevice}, devices)
}
