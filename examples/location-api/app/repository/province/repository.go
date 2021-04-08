package province_repository

import "github.com/herryg91/cdd/examples/location-api/entity"

type Repository interface {
	Get(id int) (*entity.Province, error)
	GetByIds(ids []int) (map[int]entity.Province, error)
	GetAll() ([]*entity.Province, error)
	Create(in entity.Province) (*entity.Province, error)
	Update(in entity.Province) (*entity.Province, error)
	Delete(id int) error
}
