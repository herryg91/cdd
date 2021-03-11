// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: province.proto

package tbl_province

import (
	"github.com/herryg91/cdd/examples/location-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type MysqlDatasource struct {
	db        *gorm.DB
	tableName string
}

func NewMysqlDatasource(db *gorm.DB) *MysqlDatasource {
	return &MysqlDatasource{db, "tbl_province"}
}

func (r *MysqlDatasource) GetByPrimaryKey(id int) (*entity.Province, error) {
	result := &entity.Province{}
	err := r.db.Table(r.tableName).Where("id = ?", id).Scan(&result).Error
	return result, err
}

func (r *MysqlDatasource) GetAll() ([]*entity.Province, error) {
	result := []*entity.Province{}
	err := r.db.Table(r.tableName).Find(&result).Error
	return result, err
}

func (r *MysqlDatasource) Create(in entity.Province) (*entity.Province, error) {

	in.CreatedAt = time.Now()
	in.UpdatedAt = time.Now()

	err := r.db.Table(r.tableName).Create(&in).Error
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (r *MysqlDatasource) Update(in entity.Province) (*entity.Province, error) {
	in.UpdatedAt = time.Now()
	err := r.db.Table(r.tableName).Where("id = ?", in.Id).Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (r *MysqlDatasource) Delete(id int) error {
	return r.db.Table(r.tableName).Delete(&entity.Province{}, "id = ?", id).Error
}