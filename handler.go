package main

import (
	"context"
	"os"
	"time"

	"git.sr.ht/~poldi1405/glog"
)

var (
	interrupt      = make(chan os.Signal, 1)
	shutdown       = make(chan struct{})
	serverShutdown = func() error { return nil }
	exitStatus     int
)

func handleInterrupt(interrupt chan os.Signal, origcancel context.CancelFunc, shutdown chan struct{}, originalContext context.Context) {
	sig := <-interrupt
	glog.Info("received signal ", sig, " shutting down…")

	glog.Trace("creating shutdown context with 5 seconds timeout")
	ctx, cancel := context.WithTimeout(originalContext, 5*time.Second)
	defer cancel()

	shutdownCompleted := make(chan struct{})
	go func() {
		serverShutdown()
		shutdownCompleted <- struct{}{}
	}()

	select {
	case <-shutdownCompleted:
		glog.Info("Bye!")
		shutdown <- struct{}{}
	case <-ctx.Done():
		glog.Warn("Server did not shut down within timeframe of 5 seconds. Forcing shutdown.")
		exitStatus = 1
		shutdown <- struct{}{}
	}
}
