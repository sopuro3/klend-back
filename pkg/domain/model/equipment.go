package model

type Equipment struct {
	Model
	Name        string `gorm:"not null"`
	MaxQuantity int32  `gorm:"not null"`
	Note        string `gorm:"not null"` // 備考がないなら空文字で
}

func NewEquipment(name string, maxQuantity int32, note string) *Equipment {
	return &Equipment{}
}
