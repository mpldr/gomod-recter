package main

import (
	"fmt"

	"git.sr.ht/~poldi1405/glog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"internal/data"
)

func initConfig() error {
	glog.Debug("setting up config environment")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/recter/")
	viper.AddConfigPath("/data/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("recter")

	glog.Debug("setting up default config values")
	viper.SetDefault("Directories.DataDir", defaultDataDir)
	viper.SetDefault("Directories.TemplateDir", defaultDataDir+"templates/")
	viper.SetDefault("Directories.AssetDir", defaultDataDir+"templates/assets/")
	viper.SetDefault("Domain", "my-domain.com")
	viper.SetDefault("Network.Type", "unix")
	viper.SetDefault("Network.SocketPath", "/tmp/recter.sock")
	viper.SetDefault("Network.ListenAddr", "127.0.0.1:25000")
	viper.SetDefault("Proxy.Address", "https://proxy.golang.org/")
	viper.SetDefault("Proxy.IgnoreCert", false)
	viper.SetDefault("Projects", map[string]map[string]interface{}{"example": {"name": "Example Project", "Redirect": false, "Description": "The example project is an example that shows how to add a meaningful description to your project.\n\nIf you think that explaining something with itself is a bad way of explaining a thing, feel free to submit a patch. Repetition hammers the point into your head, which is why I repeat everything I say. Having a long text is a plus because long text demonstrates better what happens if you add long text for a description.", "RootPath": "my-domain.com/example", "VCS": "git", "Repo": "https://git.sr.ht/~poldi1405/gomod-recter", "Note": map[string]interface{}{"Show": true, "Text": "This project is currently looking for a new maintainer. To apply, please reach out to me@my-domain.com", "Style": "warning"}}})

	glog.Debug("setting up FS watcher")
	viper.OnConfigChange(func(e fsnotify.Event) {
		glog.Info("config has changed.")
		glog.Debug("received event: %s", e.Op.String)

		err := viper.ReadInConfig()
		if err != nil {
			glog.Errorf("error reading in updated config: %v", err)
		}

		loadProjects()
	})

	glog.Debug("reading in config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			glog.Debug("config not found. creating…")

			err := viper.WriteConfigAs(defaultConfigPath + "config.toml")
			if err != nil {
				return fmt.Errorf("could not write config file to '%s': %w", defaultConfigPath, err)
			}
		} else {
			return fmt.Errorf("could not read config: %w", err)
		}
	}

	viper.WatchConfig()
	return nil
}

func loadProjects() {
	ps := make(map[string]*data.Project)

	glog.Debug("retrieving project list")
	projectlist := viper.GetStringMap("Projects")
	glog.Tracef("project list: %v", projectlist)

	for k := range projectlist {
		glog.Debugf("setting up project with values: Name:'%s', Desc:'%s', RootPath:'%s', VCS:'%s', Repo:'%s'", viper.GetString("Projects."+k+".Name"), viper.GetString("Projects."+k+".Description"), viper.GetString("Projects."+k+".RootPath"), viper.GetString("Projects."+k+".VCS"), viper.GetString("Projects."+k+".Repo"))
		proj := &data.Project{
			Name:        viper.GetString("Projects." + k + ".Name"),
			Description: viper.GetString("Projects." + k + ".Description"),
			RootPath:    viper.GetString("Projects." + k + ".RootPath"),
			VCS:         viper.GetString("Projects." + k + ".VCS"),
			Repo:        viper.GetString("Projects." + k + ".Repo"),
		}
		proj.GetData()

		ps[k] = proj
	}

	data.SetProjectList(ps)
}
