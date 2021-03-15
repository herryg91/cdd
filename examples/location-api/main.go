package main

import (
	"github.com/herryg91/cdd/examples/location-api/config"
	"github.com/herryg91/cdd/examples/location-api/drivers/datasource/mysql/tbl_city"
	"github.com/herryg91/cdd/examples/location-api/drivers/datasource/mysql/tbl_province"
	"github.com/herryg91/cdd/examples/location-api/drivers/handler"
	pbCity "github.com/herryg91/cdd/examples/location-api/drivers/handler/grst/city"
	pbProvince "github.com/herryg91/cdd/examples/location-api/drivers/handler/grst/province"
	crud_tbl_city "github.com/herryg91/cdd/examples/location-api/usecase/crud_tbl_city"
	crud_tbl_province "github.com/herryg91/cdd/examples/location-api/usecase/crud_tbl_province"
	search_usecase "github.com/herryg91/cdd/examples/location-api/usecase/search"
	"github.com/herryg91/cdd/grst"
	loggerInterceptor "github.com/herryg91/cdd/grst/interceptor/logger"
	recoveryInterceptor "github.com/herryg91/cdd/grst/interceptor/recovery"
	sessionInterceptor "github.com/herryg91/cdd/grst/interceptor/session"

	"github.com/herryg91/hgolib/databases/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName, cfg.DBLogMode)
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	provinceDatasource := tbl_province.NewMysqlDatasource(db)
	provinceRepo := crud_tbl_province.NewRepository(db, provinceDatasource)
	provinceUsecase := crud_tbl_province.NewUsecase(provinceRepo)
	provinceHndl := handler.NewProvinceHandler(provinceUsecase)

	cityDatasource := tbl_city.NewMysqlDatasource(db)
	cityRepo := crud_tbl_city.NewRepository(db, cityDatasource)
	cityUsecase := crud_tbl_city.NewUsecase(cityRepo)
	citySearchUsecase := search_usecase.NewUsecase(search_usecase.NewRepository(cityDatasource, provinceDatasource))
	cityHndl := handler.NewCityHandler(cityUsecase, citySearchUsecase)

	grpcRestSrv, err := grst.NewServer(cfg.GrpcPort, cfg.RestPort, true,
		grst.RegisterGRPCUnaryInterceptor("session", sessionInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("recovery", recoveryInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("log", loggerInterceptor.UnaryServerInterceptor()),
		grst.AddForwardHeaderToContext([]string{"country"}),
	)

	if err != nil {
		logrus.Panicln("Failed to Initialize GRPC-REST Server:", err)
	}

	pbProvince.RegisterProvinceGrstServer(grpcRestSrv, provinceHndl)
	pbCity.RegisterCityGrstServer(grpcRestSrv, cityHndl)
	if err := <-grpcRestSrv.ListenAndServeGrst(); err != nil {
		logrus.Panicln("Failed to Run Grpcrest Server:", err)
	}
}
