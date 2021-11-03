package worker

import "sync"

// WorkerUnion combines multiple workers into one organization that can get
// Bust()ed if needed.
type WorkerUnion struct {
	broadcaster  *sync.Cond
	broadcastMtx sync.Mutex
}

// NewUnion creates a new union
func NewUnion() *WorkerUnion {
	wo := &WorkerUnion{}
	wo.broadcaster = sync.NewCond(&wo.broadcastMtx)

	return wo
}

// Bust stops all currently active workers. This does not impact currently
// running tasks. Workers that are currently active, will complete their
// assigned function and exit afterwards.
func (wo *WorkerUnion) Bust() {
	wo.broadcaster.Broadcast()
}
