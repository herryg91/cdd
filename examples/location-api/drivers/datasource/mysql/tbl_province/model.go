package tbl_province

import "github.com/herryg91/cdd/examples/location-api/entity"

func (model *ProvinceModel) ToProvinceEntity() *entity.Province {
	return &entity.Province{
		Id:   model.Id,
		Name: model.Name,
	}
}
func (ProvinceModel) FromProvinceEntity(in entity.Province) *ProvinceModel {
	return &ProvinceModel{
		Id:   in.Id,
		Name: in.Name,
	}
}
