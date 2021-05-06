package configs

import "time"

var GlobalConfig = &Global{}

type Global struct {
	Server server
	Mysql  Mysql
	Rpc    Rpc
}

type server struct {
	AppName        string        `yaml:"app_name"`
	GRPCAddr       string        `yaml:"grpc_addr"`
	PrometheusAddr string        `yaml:"prometheus_addr"`
	Timeout        time.Duration `yaml:"timeout"`
}
