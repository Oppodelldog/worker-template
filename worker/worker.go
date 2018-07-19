package worker

import (
	"fmt"
	"time"
	"github.com/Oppodelldog/worker-template/signalcontext"
	"sync"
)

func Work(workerNo int, wg *sync.WaitGroup) {

	gracefulShutdown := signalcontext.New().Done()
	for {
		select {
		case <-gracefulShutdown:
			fmt.Printf("Graceful shutdown worker #%v\n", workerNo)
			wg.Done()
			return
		default:
			fmt.Printf("#%v working\n", workerNo)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
