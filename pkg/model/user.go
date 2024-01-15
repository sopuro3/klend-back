package model

import (
	"github.com/google/uuid"
	"github.com/sopuro3/klend-back/pkg/password"
)

type User struct {
	Model
	ExternalUserID string `gorm:"size:64;not null"` // ModelのIDと区別するため
	Email          string `gorm:"size:64;not null"`
	UserName       string `gorm:"size:64;not null"`
	HashedPassword string `gorm:"not null"` // argon2
}

func NewUser(externalUserID, email, userName string, hashedPassword password.EncodedPassword) *User {
	return &User{
		Model{
			ID: uuid.Must(uuid.NewV7()),
		},
		externalUserID,
		email,
		userName,
		string(hashedPassword),
	}
}
