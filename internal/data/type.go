package data

import (
	"sync"
	"time"

	"git.sr.ht/~poldi1405/glog"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	RootPath string `json:"importpath"`
	VCS      string `json:"vcs"`
	Repo     string `json:"repository"`
	License string `json:"license"`
	Redirect bool `json:"-"`

	Note *Note `json:"-"`

	Versions []string `json:"-"`

	LatestCommitTime time.Time `json:"last_commit_time"`
	LatestCommitHash string `json:"last_commit_hash"`
}

type Note struct {
	Show bool
	Style string
	Text string
}

var (
	projectList    map[string]*Project
	projectListMtx sync.RWMutex
)

func GetProjectList() map[string]Project {
	glog.Debug("locking mutex")
	projectListMtx.RLock()
	defer projectListMtx.RUnlock()

	glog.Debug("creating copy of map")
	cpy := make(map[string]Project)

	for k, v := range projectList {
		glog.Tracef("copying key '%s'", k)
		cpy[k] = *v
	}

	return cpy
}

func SetProjectList(projs map[string]*Project) {
	glog.Tracef("setting map: %v", projs)
	glog.Debug("creating copy of map")
	cpy := make(map[string]*Project)

	for k, v := range projs {
		glog.Tracef("copying key '%s'", k)
		cpy[k] = v
	}
	glog.Tracef("copied map: %v", cpy)

	glog.Debug("locking mutex")
	projectListMtx.Lock()
	defer projectListMtx.Unlock()

	glog.Debug("setting list")
	projectList = cpy
}
