//nolint:ireturn // domainとinfraにわけたときにはinterfaceを返す必要がある
package repository

import (
	"github.com/google/uuid"
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

	if err := er.db.Find(&equipment).Error; err != nil {
		return nil, err
	}

	return &equipment, nil
}

func (er *equipmentRepository) FindAll() ([]*model.Equipment, error) {
	var equipments []*model.Equipment

	if err := er.db.Find(&equipments).Error; err != nil {
		return nil, err
	}

	return equipments, nil
}

func (er *equipmentRepository) Create(equipment *model.Equipment) error {
	if err := er.db.Create(equipment).Error; err != nil {
		return err
	}

	return nil
}

func (er *equipmentRepository) Update(equipment *model.Equipment) error {
	if err := er.db.Save(equipment).Error; err != nil {
		return err
	}

	return nil
}

func (er *equipmentRepository) Delete(equipment *model.Equipment) error {
	if err := er.db.Delete(equipment).Error; err != nil {
		return err
	}

	return nil
}

