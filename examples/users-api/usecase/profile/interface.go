package profile_usecase

import (
	"context"

	"github.com/herryg91/cdd/examples/users-api/entity"
)

type Repository interface {
	GetProfile(ctx context.Context, id int) (*entity.UserProfile, error)
}

type UseCase interface {
	GetProfile(ctx context.Context, id int) (*entity.UserProfile, error)
}
