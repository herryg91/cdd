package graceful

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/herryg91/cdd/grst"
	"github.com/sirupsen/logrus"
)

func StartGrstServerWithGracefulShutdown(srv grst.Server, timeout_graceful int) {
	go func(server grst.Server) {
		if err := <-server.ListenAndServeGrst(); err != nil {
			logrus.Fatalln("Failed to Run GRST Server:", err)
		}
	}(srv)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout_graceful)*time.Second)
	defer cancel()
	logrus.Infof("[ %v] Grst Server is shutting down... timeout graceful: %ds.", timeout_graceful, time.Now().UTC())

	if err := srv.GetRestServer().Shutdown(ctx); err != nil {
		logrus.Fatalln("Rest Server Shutdown Failed: ", err)
	}
	srv.GetGrpcServer().GracefulStop()

	logrus.Infoln("[%v] Server is succesfully exited", time.Now().UTC())
}
