package signalcontext

import (
	"os"
	"os/signal"
	"context"
)

// New creates a context that will cancel for Interrupt
func New() context.Context {
	return NewWithSignals(os.Interrupt)
}

// NewWithSignals creates a context that will cancel for the given signals
func NewWithSignals(signals ...os.Signal) context.Context {

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, signals...)

		for {
			select {
			case <-c:
				cancelFunc()
				return
			}
		}
	}()

	return ctx
}
