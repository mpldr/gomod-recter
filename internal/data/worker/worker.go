package worker

import (
	"time"

	"git.sr.ht/~poldi1405/glog"
)

// doWorkInterval runs the given task repeatedly after every interval
func (wo *WorkerUnion) doWorkInterval(f func(), interval time.Duration) {
	for {
		select {
		case <-time.After(interval):
			f()
		case <-wo.workctx.Done():
			glog.Debug("work canceled")
			return
		}
	}
}

// AddInterval starts a worker with a given task, that will run immediately and
// once after every interval.
func (wo *WorkerUnion) AddInterval(f func(), interval time.Duration) {
	go wo.doWorkInterval(f, interval)
}
