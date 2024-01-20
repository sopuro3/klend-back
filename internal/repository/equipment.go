//nolint:ireturn  // domainとinfraにわけたときにはinterfaceを返す必要がある
package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/internal/model"
)

type EquipmentRepository interface {
	Find(id uuid.UUID) (*model.Equipment, error)
	FindAll() ([]*model.Equipment, error)
	Create(equipment *model.Equipment) error
	Update(equipment *model.Equipment) error
	Delete(equipment *model.Equipment) error
}

type equipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) EquipmentRepository {
	return &equipmentRepository{
		db: db,
	}
}

func (er *equipmentRepository) Find(id uuid.UUID) (*model.Equipment, error) {
	equipment := model.Equipment{Model: model.Model{ID: id}}

	if err := er.db.Take(&equipment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil //nolint:nilnil
		}

		return nil, err
	}

	return &equipment, nil
}

func (er *equipmentRepository) FindAll() ([]*model.Equipment, error) {
	var equipments []*model.Equipment

	result := er.db.Find(&equipments)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return equipments, nil
}

func (er *equipmentRepository) Create(equipment *model.Equipment) error {
	if err := er.db.Create(equipment).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return ErrConflict
			}
		}

		return err
	}

	return nil
}

func (er *equipmentRepository) Update(equipment *model.Equipment) error {
	// global updateを防ぐ
	if equipment.ID == (uuid.UUID{}) {
		return ErrIDIsEmpty
	}

	// 存在しない場合にInsertされるのを防ぐ
	err := er.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Take(&model.Equipment{Model: model.Model{ID: equipment.ID}}).Error; err != nil {
			return err
		}

		if err := tx.Save(equipment).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrRecodeNotFound
		}

		return err
	}

	return nil
}

func (er *equipmentRepository) Delete(equipment *model.Equipment) error {
	if equipment.ID == (uuid.UUID{}) {
		return ErrIDIsEmpty
	}

	if err := er.db.Delete(equipment).Error; err != nil {
		return err
	}

	return nil
}
