package configs

import "time"

var GlobalConfig = &Global{}

type Global struct {
	Server server
	Mysql  Mysql
	Rpc    Rpc
}

type server struct {
	AppName   string        `yaml:"app_name"`
	GRPCAddr  string        `yaml:"grpc_addr"`
	MachineId uint16        `yaml:"machine_id"`
	HttpAddr  string        `yaml:"http_addr"`
	Timeout   time.Duration `yaml:"timeout"`
}
