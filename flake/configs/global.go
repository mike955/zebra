package configs

import "time"

var GlobalConfig = &Global{}

type Global struct {
	Server server
}

type server struct {
	AppName        string        `yaml:"app_name"`
	GRPCAddr       string        `yaml:"grpc_addr"`
	MachineId      uint16        `yaml:"machine_id"`
	PrometheusAddr string        `yaml:"prometheus_addr"`
	Timeout        time.Duration `yaml:"timeout"`
}
