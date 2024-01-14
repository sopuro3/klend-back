package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/pkg/model"
)

type UserRepository interface {
	Find(id uuid.UUID) (*model.User, error)
	FindByUserID(userID string) (*model.User, error)
	FindByUserName(userName string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Find(id uuid.UUID) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByUserID(userID string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByUserName(userName string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindAll() ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Create(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Delete(user *model.User) error {
	//TODO implement me
	panic("implement me")
}
