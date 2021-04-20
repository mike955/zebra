package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/mike955/zebra/gateway/configs"
	"github.com/mike955/zebra/gateway/routers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	h "github.com/mike955/zebra/pkg/transform/http"
)

func NewHTTPServer() *h.Server {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	router := mux.NewRouter()
	routers.Route(router, logger)

	var opts = []h.ServerOption{
		h.Address(configs.GlobalConfig.Server.GRPCAddr),
		h.Timeout(configs.GlobalConfig.Server.Timeout),
		h.Router(router),
		h.Logger(logger),
	}
	srv := h.NewServer(opts...)
	return srv
}

func RunHTTPServer(server *h.Server) (err error) {
	go func() {
		if err := server.Run(); err != nil {
			server.Logger.Errorf("server start error: %s", err.Error())
			return
		}
	}()
	go http.ListenAndServe(configs.GlobalConfig.Server.PrometheusAddr, promhttp.Handler())
	handleHTTPServerSignals(server)
	return
}

func handleHTTPServerSignals(server *h.Server) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt) // stop process

	server.Logger.Info("listen quit signal ...")
	select {
	case signal := <-signalCh:
		switch signal {
		case syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt:
			server.Logger.Infof("stopping process on %s signal", fmt.Sprintf("%s", signal))
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				server.Logger.Errorf(fmt.Sprintf("quit process error|error:%s", err.Error()))
				os.Exit(1)
			}
			server.Logger.Infof(fmt.Sprintf("quit process"))
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}

func InitConfig(conf string) {
	confData, err := ioutil.ReadFile(conf)
	if err != nil {
		panic("read config file error: " + err.Error())
	}
	if err := yaml.Unmarshal(confData, configs.GlobalConfig); err != nil {
		panic("parse config file error: " + err.Error())
	}
	// data.InitSf(configs.GlobalConfig.Server.MachineId)
}
