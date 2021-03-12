package search_usecase

import (
	"fmt"

	"github.com/herryg91/cdd/examples/location-api/drivers/datasource/mysql/tbl_city"
	"github.com/herryg91/cdd/examples/location-api/drivers/datasource/mysql/tbl_province"
	"github.com/herryg91/cdd/examples/location-api/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlRepository struct {
}

type repository struct {
	cityDatasource     *tbl_city.MysqlDatasource
	provinceDatasource *tbl_province.MysqlDatasource
}

func NewRepository(cityDatasource *tbl_city.MysqlDatasource, provinceDatasource *tbl_province.MysqlDatasource) Repository {
	return &repository{cityDatasource, provinceDatasource}
}

func (r *repository) Search(keyword string) ([]entity.CityProfile, error) {
	searchResult, err := r.cityDatasource.Search(keyword)
	if err != nil {
		return []entity.CityProfile{}, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	provinceDatas, err := r.provinceDatasource.GetByIds(extractCityIds(searchResult))
	if err != nil {
		return []entity.CityProfile{}, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	result := []entity.CityProfile{}
	for _, s := range searchResult {
		provinceName := ""
		if p, ok := provinceDatas[s.ProvinceId]; ok {
			provinceName = p.Name
		}
		result = append(result, entity.CityProfile{}.FromCity(*s, provinceName))
	}
	return result, nil
}

func extractCityIds(arr []*entity.City) []int {
	distinctId := map[int]bool{}
	result := []int{}
	for _, a := range arr {
		if _, ok := distinctId[a.ProvinceId]; !ok {
			result = append(result, a.ProvinceId)
			distinctId[a.ProvinceId] = true
		}
	}
	return result
}
