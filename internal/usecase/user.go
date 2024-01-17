package usecase

import "github.com/sopuro3/klend-back/internal/repository"

type UserUseCase struct {
	ur repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		ur: ur,
	}
}

func (uu *UserUseCase) CreateUser(externalUserID, email, userName, password string) error {
	_ = externalUserID
	_ = email
	_ = userName
	_ = password

	return nil
}
