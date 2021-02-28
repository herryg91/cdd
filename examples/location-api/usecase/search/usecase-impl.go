package search_usecase

import (
	"github.com/herryg91/cdd/examples/location-api/entity"
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

func (uc *usecase) Search(keyword string) ([]entity.CityProfile, error) {
	return uc.repo.Search(keyword)
}
