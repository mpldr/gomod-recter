package worker

import (
	"context"
)

// WorkerUnion combines multiple workers into one organization that can get
// Bust()ed if needed.
type WorkerUnion struct {
	workctx    context.Context
	workcancel context.CancelFunc
}

// NewUnion creates a new union
func NewUnion() *WorkerUnion {
	wo := &WorkerUnion{}
	wo.workctx, wo.workcancel = context.WithCancel(context.Background())

	return wo
}

// Bust stops all currently active workers. This does not impact currently
// running tasks. Workers that are currently active, will complete their
// assigned function and exit afterwards.
func (wo *WorkerUnion) Bust() {
	wo.workcancel()
}
