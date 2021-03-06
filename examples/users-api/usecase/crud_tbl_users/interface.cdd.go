// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: users.proto

package crud_tbl_users

import (
	"github.com/herryg91/cdd/examples/users-api/entity"
)

type Repository interface {
	additional_repository
	GetByPrimaryKey(id int) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Create(in entity.User) (*entity.User, error)
	Update(in entity.User) (*entity.User, error)
	Delete(id int) error
}

type UseCase interface {
	additional_usecase
	GetByPrimaryKey(id int) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Create(in entity.User) (*entity.User, error)
	Update(in entity.User) (*entity.User, error)
	Delete(id int) error
}

// Please write code below in interface.go
/*
	package crud_tbl_users

	type additional_repository interface{
		// AdditionalFunc1()
		// AdditionalFunc2()
	}

	type additional_usecase interface {
		// AdditionalFunc1()
		// AdditionalFunc2()
	}
*/
