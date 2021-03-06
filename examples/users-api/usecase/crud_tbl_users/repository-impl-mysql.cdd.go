// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: users.proto
package crud_tbl_users

import (
	"fmt"
	tbl_users_ds "github.com/herryg91/cdd/examples/users-api/drivers/datasource/mysql/tbl_users"
	"github.com/herryg91/cdd/examples/users-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type repository struct {
	db *gorm.DB
	ds *tbl_users_ds.MysqlDatasource
}

func NewRepository(db *gorm.DB, ds *tbl_users_ds.MysqlDatasource) Repository {
	return &repository{db, ds}
}
func (r *repository) GetByPrimaryKey(id int) (*entity.User, error) {
	out, err := r.ds.GetByPrimaryKey(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out, err
}
func (r *repository) GetAll() ([]*entity.User, error) {
	out, err := r.ds.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out, err
}
func (r *repository) Create(in entity.User) (*entity.User, error) {
	out, err := r.ds.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out, err
}
func (r *repository) Update(in entity.User) (*entity.User, error) {
	out, err := r.ds.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out, err
}
func (r *repository) Delete(id int) error {
	err := r.ds.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return err
}
