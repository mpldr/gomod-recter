package main

import (
	"net"
	"os"

	"mpldr.codes/recter/internal/handler"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

var Version = "devel"

func main() {
	glog.SetLevel(glog.INFO)
	glog.Info("starting up recter version " + Version)
	glog.Debug("setting up config")
	initConfig()

	loadProjects()

	var address string

	switch viper.GetString("Network.Type") {
	case "tcp", "tcp4", "tcp6":
		address = viper.GetString("Network.ListenAddr")
	case "unix":
		address = viper.GetString("Network.SocketPath")
		defer os.Remove(address)
	default:
		glog.Fatalf("invalid network type '%s'. Allowed types are tcp{,4,6} and unix", viper.GetString("Network.Type"))
		os.Exit(1)
	}

	s := &fasthttp.Server{
		Handler:                      handler.FasthttpHandler,
		Name:                         "gomod-recter/" + Version,
		Concurrency:                  128,
		GetOnly:                      false,
		DisablePreParseMultipartForm: true,
		LogAllErrors:                 true,
		SecureErrorLogMessage:        true,
		CloseOnShutdown:              true,
		Logger:                       &logtype{},
	}
	serverShutdown = s.Shutdown

	lis, err := net.Listen(viper.GetString("Network.Type"), address)
	if err != nil {
		glog.Fatalf("could not listen on %s:%s", viper.GetString("Network.Type"), address)
		os.Exit(1)
	}
	defer lis.Close()

	glog.Infof("serving on %s:%s", viper.GetString("Network.Type"), address)
	err = s.Serve(lis)
	if err != nil {
		glog.Fatalf("error serving webserver: %v", err)
		return
	}

	<-shutdown
	os.Exit(exitStatus)
}

type logtype struct{}

func (l *logtype) Printf(fmt string, args ...interface{}) {
	glog.Debugf(fmt, args...)
}
