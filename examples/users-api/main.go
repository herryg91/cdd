package main

import (
	"github.com/herryg91/cdd/examples/users-api/config"
	"github.com/herryg91/cdd/examples/users-api/drivers/datasource/mysql/tbl_users"
	pbProvince "github.com/herryg91/cdd/examples/users-api/drivers/external/grst/province"
	"github.com/herryg91/cdd/examples/users-api/drivers/handler"
	pbUsers "github.com/herryg91/cdd/examples/users-api/drivers/handler/grst/users"
	profile_usecase "github.com/herryg91/cdd/examples/users-api/usecase/profile"
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

	usersDatasource := tbl_users.NewMysqlDatasource(db)
	provinceClient, err := pbProvince.NewProvinceGrstClient(cfg.LocationApi, nil)
	if err != nil {
		panic(err)
	}
	profileRepo := profile_usecase.NewRepository(db, usersDatasource, provinceClient)
	profileUsecase := profile_usecase.NewUsecase(profileRepo)
	usersHndl := handler.NewHandler(profileUsecase)

	grpcRestSrv, err := grst.NewServer(cfg.GrpcPort, cfg.RestPort, true,
		grst.RegisterGRPCUnaryInterceptor("session", sessionInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("recovery", recoveryInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("log", loggerInterceptor.UnaryServerInterceptor()),
		grst.AddForwardHeaderToContext([]string{"country"}),
	)

	if err != nil {
		logrus.Panicln("Failed to Initialize GRPC-REST Server:", err)
	}

	pbUsers.RegisterUsersGrstServer(grpcRestSrv, usersHndl)
	if err := <-grpcRestSrv.ListenAndServeGrst(); err != nil {
		logrus.Panicln("Failed to Run Grpcrest Server:", err)
	}
}
