package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"mpldr.codes/recter/internal/data"
	"mpldr.codes/recter/internal/data/worker"

	"git.sr.ht/~poldi1405/glog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func initConfig() error {
	glog.Debug("setting up config environment")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(defaultConfigPath)
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("recter")

	glog.Debug("setting up default config values")
	viper.SetDefault("Directories.DataDir", defaultDataDir)
	viper.SetDefault("Directories.TemplateDir", defaultTemplateDir+"templates/")
	viper.SetDefault("Directories.AssetDir", defaultTemplateDir+"templates/assets/")
	viper.SetDefault("Domain", "my-domain.com")
	viper.SetDefault("VersionRefreshInterval", 5*time.Minute)
	viper.SetDefault("Network.Type", defaultNetwork)
	viper.SetDefault("Network.SocketPath", "/tmp/recter.sock")
	viper.SetDefault("Network.ListenAddr", defaultListenAddr)
	viper.SetDefault("Proxy.Address", "https://proxy.golang.org/")
	viper.SetDefault("Proxy.IgnoreCert", false)
	viper.SetDefault("Projects", map[string]map[string]interface{}{
		"example": {
			"Name":          "Example Project",
			"Redirect":      false,
			"Description":   "The example project is an example that shows how to add a meaningful description to your project.\n\nIf you think that explaining something with itself is a bad way of explaining a thing, feel free to submit a patch. Repetition hammers the point into your head, which is why I repeat everything I say. Having a long text is a plus because long text demonstrates better what happens if you add long text for a description.",
			"VCS":           "git",
			"Repo":          "https://git.sr.ht/~poldi1405/gomod-recter",
			"License":       "MIT",
			"DefaultBranch": "master",
			"GoSourceFmt":   "",
			"Note": map[string]interface{}{
				"Show":  true,
				"Text":  "This project is currently looking for a new maintainer. To apply, please reach out to me@my-domain.com",
				"Style": "warning",
			},
		},
	})

	glog.Debug("setting up FS watcher")
	viper.OnConfigChange(func(e fsnotify.Event) {
		glog.Info("config has changed.")
		glog.Debugf("received event: %s", e.Op.String())

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
			if errors.Is(err, os.ErrPermission) {
				err = viper.WriteConfigAs("." + string(os.PathSeparator) + "config.toml")
			}
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

var union = worker.NewUnion()

func loadProjects() {
	ps := make(map[string]*data.Project)
	union.Bust()

	glog.Debug("retrieving project list")
	projectlist := viper.GetStringMap("Projects")
	glog.Tracef("project list: %v", projectlist)

	union = worker.NewUnion()

	var wg sync.WaitGroup
	for k := range projectlist {
		wg.Add(1)

		redirTo := viper.GetString("Projects." + k + ".RedirectTo")
		if redirTo == "" {
			redirTo = viper.GetString("Projects." + k + ".Repo")
		}

		glog.Debugf("Project data: %+v", viper.Get("Projects."+k))

		proj := &data.Project{
			Key:           k,
			Name:          viper.GetString("Projects." + k + ".Name"),
			Description:   viper.GetString("Projects." + k + ".Description"),
			RootPath:      viper.GetString("Domain") + "/" + k,
			VCS:           viper.GetString("Projects." + k + ".VCS"),
			Repo:          viper.GetString("Projects." + k + ".Repo"),
			License:       viper.GetString("Projects." + k + ".License"),
			Redirect:      viper.GetBool("Projects." + k + ".Redirect"),
			RedirectTo:    redirTo,
			Hidden:        viper.GetBool("Projects." + k + ".Hidden"),
			DefaultBranch: viper.GetString("Projects." + k + ".DefaultBranch"),
			GoSourceFmt:   viper.GetString("Projects." + k + ".GoSourceFmt"),
			Note: &data.Note{
				Show:  viper.GetBool("Projects." + k + ".Note.Show"),
				Style: viper.GetString("Projects." + k + ".Note.Style"),
				Text:  viper.GetString("Projects." + k + ".Note.Text"),
			},
			Versions:         make([]string, 0),
			LatestCommitTime: time.Time{},
			LatestCommitHash: "",
		}
		glog.Debugf("Parsed project: %+v", proj)
		go func() {
			defer wg.Done()
			proj.GetData()
			union.AddInterval(func() { proj.GetData() }, viper.GetDuration("VersionRefreshInterval"))
		}()

		ps[k] = proj
	}
	wg.Wait()

	data.SetProjectList(ps)
}
