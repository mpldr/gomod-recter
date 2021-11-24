//go:build !docker

package main

var (
	defaultDataDir     = "/opt/recter/"
	defaultTemplateDir = defaultDataDir + "themes/"
	defaultConfigPath  = "/etc/recter/"
)
