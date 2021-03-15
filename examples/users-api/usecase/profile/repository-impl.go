package profile_usecase

import (
	"context"
	"fmt"

	tbl_users_ds "github.com/herryg91/cdd/examples/users-api/drivers/datasource/mysql/tbl_users"
	pbProvince "github.com/herryg91/cdd/examples/users-api/drivers/external/grst/province"
	"github.com/herryg91/cdd/examples/users-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type repository struct {
	db             *gorm.DB
	ds             *tbl_users_ds.MysqlDatasource
	provinceClient pbProvince.ProvinceClient
}

func NewRepository(db *gorm.DB, ds *tbl_users_ds.MysqlDatasource, provinceClient pbProvince.ProvinceClient) Repository {
	return &repository{db, ds, provinceClient}
}
func (r *repository) GetProfile(ctx context.Context, id int) (*entity.UserProfile, error) {
	user, err := r.ds.GetByPrimaryKey(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	provinceResp, err := r.provinceClient.Get(ctx, &pbProvince.GetReq{
		Id: int32(id),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %s - %s", ErrClient, "location-api (province)", err.Error())
	}

	out := entity.UserProfile{}.FromUser(*user.ToUserEntity(), provinceResp.Name)
	return &out, err
}
