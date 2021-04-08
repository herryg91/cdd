package city_repository

import "github.com/herryg91/cdd/examples/location-api/entity"

type Repository interface {
	GetById(id int) (*entity.City, error)
	GetAll() ([]*entity.City, error)
	Create(in entity.City) (*entity.City, error)
	Update(in entity.City) (*entity.City, error)
	Delete(id int) error
	Search(keyword string) ([]*entity.City, error)
}
