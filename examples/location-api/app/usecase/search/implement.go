package search_usecase

import (
	"fmt"

	city_repository "github.com/herryg91/cdd/examples/location-api/app/repository/city"
	province_repository "github.com/herryg91/cdd/examples/location-api/app/repository/province"
	"github.com/herryg91/cdd/examples/location-api/entity"
)

type usecase struct {
	city_repo     city_repository.Repository
	province_repo province_repository.Repository
}

func New(city_repo city_repository.Repository, province_repo province_repository.Repository) UseCase {
	return &usecase{
		city_repo:     city_repo,
		province_repo: province_repo,
	}
}

func (uc *usecase) Search(keyword string) ([]entity.CityProfile, error) {
	datas, err := uc.city_repo.Search(keyword)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	mapOfProvinces, err := uc.province_repo.GetByIds(entity.Cities(datas).GetProvinceIds())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	resp := []entity.CityProfile{}
	for _, data := range datas {
		provinceName := ""
		if val, ok := mapOfProvinces[data.ProvinceId]; ok {
			provinceName = val.Name
		}
		resp = append(resp, entity.CityProfile{}.FromCity(*data, provinceName))
	}
	return resp, nil
}
