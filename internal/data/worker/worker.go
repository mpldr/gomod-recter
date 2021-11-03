package worker

import "time"

// doWorkInterval runs the given task repeatedly after every interval
func (wo *WorkerUnion) doWorkInterval(f func(), interval time.Duration) {
	stop := make(chan struct{})
	go func() {
		wo.broadcaster.Wait()
		stop <- struct{}{}
	}()

	for {
		select {
		case <-time.After(interval):
			f()
		case <-stop:
			return
		}
	}
}

// AddInterval starts a worker with a given task, that will run immediately and
// once after every interval.
func (wo *WorkerUnion) AddInterval(f func(), interval time.Duration) {
	go wo.doWorkInterval(f, interval)
}
