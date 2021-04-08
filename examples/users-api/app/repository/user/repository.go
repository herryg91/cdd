package user_repository

import "github.com/herryg91/cdd/examples/users-api/entity"

type Repository interface {
	GetById(id int) (*entity.User, error)
	GetProfileById(id int) (*entity.UserProfile, error)
}
