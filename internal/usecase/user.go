package usecase

import "github.com/sopuro3/klend-back/internal/repository"

type UserUseCase struct {
	r repository.BaseRepository
}

func NewUserUseCase(r repository.BaseRepository) *UserUseCase {
	return &UserUseCase{
		r: r,
	}
}

func (uu *UserUseCase) CreateUser(externalUserID, email, userName, password string) error {
	_ = externalUserID
	_ = email
	_ = userName
	_ = password

	return nil
}
