package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/mike955/zebra/account/internal/server"
)

var (
	version bool
	conf    string

	BuildTime       = ""
	GitCommitID     = ""
	GitCommitBranch = ""
	GoVersion       = runtime.Version()
)

func init() {
	flag.StringVar(&conf, "f", "", "-f <config>")
	flag.BoolVar(&version, "v", false, "-v")
	flag.Parse()
	if conf == "" {
		panic("not found config file, use: -f config.yaml")
	}
	server.InitConfig(conf)
}

func main() {
	// flag.Parse()
	if version == true {
		fmt.Println("BuildTime: ", BuildTime)
		fmt.Println("GitCommitID: ", GitCommitID)
		fmt.Println("GitCommitBranch: ", GitCommitBranch)
		fmt.Println("GoVersion: ", GoVersion)
		fmt.Println("GitCommitID: ", BuildTime)
		return
	}
	// if conf == "" {
	// 	panic("not found config file, use: -f config.yaml")
	// }

	grpcServe := server.NewGRPCServer(conf)
	if err := server.RunGRPCServer(grpcServe); err != nil {
		grpcServe.Logger.Errorf("grpc server run error:%s", err.Error())
	}
}
