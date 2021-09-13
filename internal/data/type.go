package data

import (
	"sync"
	"time"

	"git.sr.ht/~poldi1405/glog"
)

type Project struct {
	Name        string
	Description string

	RootPath string
	VCS      string
	Repo     string
	Redirect bool

	Versions []string

	LatestCommitTime time.Time
	LatestCommitHash string
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
