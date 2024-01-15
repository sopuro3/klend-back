package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/pkg/model"
)

type UserRepository interface {
	Find(id uuid.UUID) (*model.User, error)
	FindByUserID(userID string) (*model.User, error)
	FindByUserName(userName string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindExistUser(user *model.User) (bool, error)
	FindAll() ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Find(id uuid.UUID) (*model.User, error) {
	user := model.User{Model: model.Model{ID: id}}
	if err := u.db.Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindByUserID(externalUserID string) (*model.User, error) {
	user := model.User{}
	if err := u.db.Where(&model.User{ExternalUserID: externalUserID}).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindByUserName(userName string) (*model.User, error) {
	user := model.User{}
	if err := u.db.Where(&model.User{UserName: userName}).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := u.db.Where(&model.User{Email: email}).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return users, nil
}

func (u *userRepository) FindExistUser(user *model.User) (bool, error) {
	if user.ExternalUserID == "" && user.Email == "" {
		return false, ErrEmptyData
	}
	var count int64
	if err := u.db.Model(&model.User{}).Or(&model.User{ExternalUserID: user.ExternalUserID}).Or(&model.User{Email: user.Email}).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
func (u *userRepository) Create(user *model.User) error {
	if user.ID == (uuid.UUID{}) || user.ExternalUserID == "" || user.Email == "" || user.UserName == "" || user.HashedPassword == "" {
		return ErrEmptyData
	}
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user *model.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(user *model.User) error {
	if user.ID == (uuid.UUID{}) {
		return ErrIDIsEmpty
	}
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
