package tbl_city

import (
	"github.com/herryg91/cdd/examples/location-api/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (r *MysqlDatasource) Search(keyword string) ([]*entity.City, error) {
	result := []*entity.City{}
	err := r.db.Table(r.tableName).Where("name like ?", "%"+keyword+"%").Find(&result).Error
	return result, err
}
