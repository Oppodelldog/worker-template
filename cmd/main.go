package main

import (
	"time"
	"syscall"
	"github.com/Oppodelldog/worker-template/worker"
	"sync"
)

func main() {

	go initSelfShutdown() // yes indeed just for simulation purpose

	// spawn some workers
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)

		go worker.Work(i, wg)

	}
	wg.Wait() // synchronized by a simple WaitGroup
}

func initSelfShutdown() {
	selfShutdownTimer := time.NewTimer(time.Second * 4)
	for {
		select {
		case <-selfShutdownTimer.C:
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}
	}
}
