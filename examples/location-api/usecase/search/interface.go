package search_usecase

import "github.com/herryg91/cdd/examples/location-api/entity"

type Repository interface {
	Search(keyword string) ([]entity.CityProfile, error)
}

type UseCase interface {
	Search(keyword string) ([]entity.CityProfile, error)
}
