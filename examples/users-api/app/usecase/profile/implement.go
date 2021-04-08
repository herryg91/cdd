package profile

import (
	"fmt"

	user_repository "github.com/herryg91/cdd/examples/users-api/app/repository/user"
	"github.com/herryg91/cdd/examples/users-api/entity"
)

type usecase struct {
	user_repo user_repository.Repository
}

func New(user_repo user_repository.Repository) UseCase {
	return &usecase{user_repo: user_repo}
}

func (uc *usecase) GetProfile(id int) (*entity.UserProfile, error) {
	data, err := uc.user_repo.GetProfileById(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
