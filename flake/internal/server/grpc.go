package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	configs "github.com/mike955/zebra/flake/configs"
	"github.com/mike955/zebra/flake/internal/data"
	"github.com/mike955/zebra/flake/internal/service"
	"github.com/mike955/zebra/pkg/transform/grpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"

	pb "github.com/mike955/zebra/api/flake"
)

// func _init(conf string) {
// 	initConfig(conf)
// 	// initRpcClient()
// }

func NewGRPCServer(conf string) (server *grpc.Server) {
	// _init(conf)
	logger := newLogger()

	config := configs.GlobalConfig.Server
	var opts = []grpc.ServerOption{
		grpc.Logger(logger),
		grpc.Address(config.GRPCAddr),
		grpc.Timeout(config.Timeout),
		grpc.GrpcUnaryServerInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.GrpcDefaultUnaryServerInterceptor(),
	}

	server = grpc.NewServer("flake", opts...)
	log := server.Logger.WithFields(logrus.Fields{"app": "flake"})
	s := service.NewFlakeService(log)
	pb.RegisterFlakeServiceServer(server, s)
	reflection.Register(server.Server) // Register reflection service on gRPC server.
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(server.Server)
	grpc.GrpcHealthCheck(server.Server)
	return
}

func RunGRPCServer(server *grpc.Server) (err error) {
	go func() {
		if err := server.Start(); err != nil {
			server.Logger.Errorf("server start error: %s", err.Error())
			return
		}
	}()
	go http.ListenAndServe(configs.GlobalConfig.Server.PrometheusAddr, promhttp.Handler())
	handleGRPCServerSignals(server)
	return
}

func handleGRPCServerSignals(server *grpc.Server) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt) // stop process

	server.Logger.Info("listen quit signal ...")
	select {
	case signal := <-signalCh:
		switch signal {
		case syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt:
			server.Logger.Infof("stopping process on %s signal", fmt.Sprintf("%s", signal))
			if err := server.Stop(); err != nil {
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
	data.InitSf(configs.GlobalConfig.Server.MachineId)
}

func newLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Out = os.Stderr
	return logger
}
