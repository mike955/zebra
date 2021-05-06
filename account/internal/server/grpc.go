package server

import (
	"io/ioutil"

	_ "net/http/pprof"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	configs "github.com/mike955/zebra/account/configs"
	"github.com/mike955/zebra/account/internal/dao"
	"github.com/mike955/zebra/account/internal/service"
	pb "github.com/mike955/zebra/api/account"
	"github.com/mike955/zebra/pkg/transform/grpc"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewGRPCServer(conf string) (server *grpc.Server) {
	config := configs.GlobalConfig.Server
	var opts = []grpc.ServerOption{
		grpc.Address(config.GRPCAddr),
		grpc.Timeout(config.Timeout),
		grpc.GrpcUnaryServerInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.GrpcDefaultUnaryServerInterceptor(),

		grpc.Prometheus(true, configs.GlobalConfig.Server.PrometheusAddr),
		grpc.Reflection(),
		grpc.HealthCheck(),
	}

	server = grpc.NewServer(config.AppName, opts...)
	log := server.Logger.WithFields(logrus.Fields{"app": config.AppName})
	s := service.NewAccountService(log)
	pb.RegisterAccountServiceServer(server, s)
	return
}

func RunGRPCServer(server *grpc.Server) (err error) {
	err = server.Start()
	if err != nil {
		server.Logger.Errorf("server start error: %s", err.Error())
	}
	return
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
	dao.Init(configs.GlobalConfig.Mysql)
}
