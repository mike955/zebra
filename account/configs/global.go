package account

import "time"

var GlobalConfig = &Global{}

type Global struct {
	Server server
	Mysql  Mysql
}

type server struct {
	GRPCAddr       string        `yaml:"grpc_addr"`
	MachineId      uint16        `yaml:"machine_id"`
	PrometheusAddr string        `yaml:"prometheus_addr"`
	Timeout        time.Duration `yaml:"timeout"`
}
