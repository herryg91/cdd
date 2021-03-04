package profile_usecase

import (
	"context"

	"github.com/herryg91/cdd/examples/users-api/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) UseCase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetProfile(ctx context.Context, id int) (*entity.UserProfile, error) {
	return uc.repo.GetProfile(ctx, id)
}
