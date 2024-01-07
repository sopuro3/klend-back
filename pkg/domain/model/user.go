package model

type User struct {
	Model
	ExternalUserID string `gorm:"size:64;not null"` // ModelのIDと区別するため
	Email          string `gorm:"size:64;not null"`
	Username       string `gorm:"size:64;not null"`
	HashedPassword string `gorm:"not null"` // argon2
}
