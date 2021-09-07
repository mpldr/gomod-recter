package main

import (
	"fmt"

	"git.sr.ht/~poldi1405/glog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func initConfig() error {
	glog.Debug("setting up config environment")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/recter/")
	viper.AddConfigPath("/data/")
	viper.SetEnvPrefix("recter")

	glog.Debug("setting up default config values")
	viper.SetDefault("Directories.DataDir", defaultDataDir)
	viper.SetDefault("Directories.TemplateDir", defaultDataDir+"templates/")
	viper.SetDefault("Directories.AssetDir", defaultDataDir+"templates/assets/")
	viper.SetDefault("Domain", "my-domain.com")
	viper.SetDefault("Proxy.Address", "https://proxy.golang.org/")
	viper.SetDefault("Proxy.Insecure", false)

	viper.OnConfigChange(func(e fsnotify.Event) {
		glog.Info("config has changed.")
		glog.Debug("received event: %s", e.Op.String)

		err := viper.ReadInConfig()
		if err != nil {
			glog.Errorf("error reading in updated config: %v", err)
		}
	})
	viper.WatchConfig()

	glog.Debug("reading in config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			glog.Debug("config not found. creating…")

			err := viper.WriteConfigAs(defaultConfigPath + "config.tom")
			if err != nil {
				return fmt.Errorf("could not write config file to '%s': %w", defaultConfigPath, err)
			}
		} else {
			return fmt.Errorf("could not read config: %w", err)
		}
	}

	return nil
}
