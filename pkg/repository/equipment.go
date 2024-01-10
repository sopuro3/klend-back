package repository

import (
	"github.com/google/uuid"

	"github.com/sopuro3/klend-back/pkg/model"
)

type EquipmentRepository interface {
	Find(id uuid.UUID) (*model.Equipment, error)
	FindAll() ([]*model.Equipment, error)
	Create(equipment *model.Equipment) error
	Update(equipment *model.Equipment) error
	Delete(equipment *model.Equipment) error
}
