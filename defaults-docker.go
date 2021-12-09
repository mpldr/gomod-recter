//go:build docker

package main

var (
	defaultDataDir     = "/data/"
	defaultTemplateDir = "/themes/"
	defaultConfigPath  = "/data/"
	defaultNetwork     = "tcp"
	defaultListenAddr  = "0.0.0.0:25000"
)
