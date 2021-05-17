package server

import (
	"io/ioutil"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	configs "github.com/mike955/zebra/email/configs"
	"github.com/mike955/zebra/email/internal/dao"
	"github.com/mike955/zebra/email/internal/service"
	"github.com/mike955/zrpc/transform/grpc"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	_ "net/http/pprof"

	pb "github.com/mike955/zebra/api/email"
)

func NewGRPCServer() (server *grpc.Server) {
	config := configs.GlobalConfig.Server
	var opts = []grpc.ServerOption{
		grpc.Address(config.GRPCAddr),
		grpc.Timeout(config.Timeout),
		grpc.GrpcUnaryServerInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.GrpcDefaultUnaryServerInterceptor(),

		grpc.Prometheus(true, configs.GlobalConfig.Server.HttpAddr),
		grpc.Reflection(),
		grpc.HealthCheck(),
	}

	server = grpc.NewServer(config.AppName, opts...)
	log := server.Logger.WithFields(logrus.Fields{"app": config.AppName})
	s := service.NewEmailService(log)
	pb.RegisterEmailServiceServer(server, s)
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
	dao.Init(configs.GlobalConfig.Mysql)
}
