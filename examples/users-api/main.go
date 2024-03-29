package main

import (
	"time"

	profile_usecase "github.com/herryg91/cdd/examples/users-api/app/usecase/profile"
	pbProvince "github.com/herryg91/cdd/examples/users-api/clients/grst/province"
	"github.com/herryg91/cdd/examples/users-api/config"
	"github.com/herryg91/cdd/examples/users-api/handler"
	pbUsers "github.com/herryg91/cdd/examples/users-api/handler/grst/users"
	"github.com/herryg91/cdd/examples/users-api/pkg/conn/mysql"
	"github.com/herryg91/cdd/examples/users-api/repository/user_repository_impl"
	"github.com/herryg91/cdd/grst"
	loggerInterceptor "github.com/herryg91/cdd/grst/interceptor/logger"
	recoveryInterceptor "github.com/herryg91/cdd/grst/interceptor/recovery"
	sessionInterceptor "github.com/herryg91/cdd/grst/interceptor/session"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

func main() {
	cfg := config.New()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName,
		mysql.SetPrintLog(cfg.DBLogEnable, logger.LogLevel(cfg.DBLogLevel), time.Duration(cfg.DBLogThreshold)*time.Millisecond))
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	provinceClient, err := pbProvince.NewProvinceGrstClient(cfg.LocationApi, nil)
	if err != nil {
		panic(err)
	}
	userRepo := user_repository_impl.New(db, provinceClient)
	profileUsecase := profile_usecase.New(userRepo)
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
