package tbl_users

import "github.com/herryg91/cdd/examples/users-api/entity"

func (model *UserModel) ToUserEntity() *entity.User {
	return &entity.User{
		Id:         model.Id,
		Name:       model.Name,
		ProvinceId: model.ProvinceId,
	}
}
func (UserModel) FromUserEntity(in entity.User) *UserModel {
	return &UserModel{
		Id:         in.Id,
		Name:       in.Name,
		ProvinceId: in.ProvinceId,
	}
}
