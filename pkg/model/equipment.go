package model

import "github.com/google/uuid"

type Equipment struct {
	Model
	Name        string `gorm:"not null"`
	MaxQuantity int32  `gorm:"not null"`
	Note        string `gorm:"not null"` // 備考がないなら空文字で
}

func NewEquipment(name string, maxQuantity int32, note string) *Equipment {
	return &Equipment{
		Model: Model{
			ID: uuid.Must(uuid.NewV7()),
		},
		Name:        name,
		MaxQuantity: maxQuantity,
		Note:        note,
	}
}
