//go:build !docker

package main

var (
	defaultConfigPath = "./"
	// defaultConfigPath = "/etc/recter/"
	defaultDataDir = "/opt/recter/"
)
