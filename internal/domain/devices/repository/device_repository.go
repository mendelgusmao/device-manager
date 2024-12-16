package repository

import (
	"context"

	"github.com/jinzhu/copier"
	domainErrors "github.com/mendelgusmao/device-manager/internal/domain/devices/errors"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/models"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}

func (r *DeviceRepository) Insert(ctx context.Context, device models.Device) (*models.Device, error) {
	tx := r.db.Begin()

	if err := tx.Create(&device).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepository) FetchOne(ctx context.Context, query *models.DeviceQuery) (*models.Device, error) {
	devices, err := r.fetch(ctx, query, 1)

	if err != nil {
		return nil, err
	}

	if len(devices) == 0 {
		return nil, domainErrors.ErrorRecordNotFound
	}

	return &devices[0], nil
}

func (r *DeviceRepository) FetchMany(ctx context.Context, query *models.DeviceQuery) ([]models.Device, error) {
	devices, err := r.fetch(ctx, query, 0)

	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *DeviceRepository) Update(ctx context.Context, device models.Device) (*models.Device, error) {
	fetchedDevice, err := r.FetchOne(ctx, &models.DeviceQuery{ID: &device.ID})

	if err != nil {
		return nil, err
	}

	copier.Copy(&fetchedDevice, &device)
	tx := r.db.WithContext(ctx).Begin()
	result := tx.Save(&device)

	if err := result.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepository) Delete(ctx context.Context, query models.DeviceQuery) error {
	tx := r.db.WithContext(ctx).Begin()
	result := tx.Delete(&models.Device{ID: *query.ID})

	if err := result.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrorRecordNotFound
	}

	return nil
}

func (r *DeviceRepository) fetch(ctx context.Context, query *models.DeviceQuery, limit int) (devices []models.Device, err error) {
	tx := r.db.WithContext(ctx)

	if query != nil {
		tx = tx.Where(&query)
	}

	if limit > 0 {
		tx = tx.Limit(limit)
	}

	if err = tx.Find(&devices).Error; err != nil {
		return nil, err
	}

	return
}

func (r *DeviceRepository) Setup() error {
	return r.db.AutoMigrate(&models.Device{})
}
