//go:build !docker

package main

var (
	defaultDataDir     = "/opt/recter/"
	defaultTemplateDir = defaultDataDir + "themes/"
	defaultConfigPath  = "/etc/recter/"
	defaultNetwork     = "unix"
	defaultListenAddr  = "127.0.0.1:25000"
)
