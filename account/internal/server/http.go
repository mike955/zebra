package server

import (
	"log"
	"net/http"
	"time"

	configs "github.com/mike955/zebra/account/configs"
	"github.com/mike955/zebra/account/internal/service"
	pb "github.com/mike955/zebra/api/account"
	"github.com/sirupsen/logrus"
)

func NewHTTPServer() (handler http.Handler) {
	config := configs.GlobalConfig.Server
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log := logger.WithFields(logrus.Fields{"app": config.AppName})
	s := service.NewAccountService(log)
	handler = pb.NewAccountServiceHTTPServer(s, log)
	return
}

func RunHTTPServer(handler http.Handler) (err error) {
	config := configs.GlobalConfig.Server
	srv := &http.Server{
		Handler: handler,
		Addr:    config.HttpAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	return
}
