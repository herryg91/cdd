package tbl_city

import "github.com/herryg91/cdd/examples/location-api/entity"

func (model *CityModel) ToCityEntity() *entity.City {
	return &entity.City{
		Id:         model.Id,
		ProvinceId: model.ProvinceId,
		Name:       model.Name,
	}
}
func (CityModel) FromCityEntity(in entity.City) *CityModel {
	return &CityModel{
		Id:         in.Id,
		ProvinceId: in.ProvinceId,
		Name:       in.Name,
	}
}
