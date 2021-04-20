package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/mike955/zebra/gateway/server"
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
	if version == true {
		fmt.Println("BuildTime: ", BuildTime)
		fmt.Println("GitCommitID: ", GitCommitID)
		fmt.Println("GitCommitBranch: ", GitCommitBranch)
		fmt.Println("GoVersion: ", GoVersion)
		fmt.Println("GitCommitID: ", BuildTime)
		return
	}

	srv := server.NewHTTPServer()
	if err := server.RunHTTPServer(srv); err != nil {
		srv.Logger.Errorf("http server run error:%s", err.Error())
	}
}
