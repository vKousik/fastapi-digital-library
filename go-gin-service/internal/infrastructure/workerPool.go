package infrastructure

import (
	"log"
	"time"

	"github.com/vKousik/fastapi-digital-library/internal/domain"
)

type Job struct {
	Type    string
	Payload domain.Book
}

var JobQueue = make(chan Job, 100)

func StartWorkerPool(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			log.Printf("[Worker %d] ready", workerID)
			for job := range JobQueue {
				process(workerID, job)
			}
		}(i)
	}
}

func process(workerID int, job Job) {
	log.Printf(
		"[Worker %d] started job: %s for %s",
		workerID,
		job.Type,
		job.Payload.Title,
	)
	time.Sleep(4 * time.Second)
	log.Printf(
		"[Worker %d] finished job: %s",
		workerID,
		job.Type,
	)
}
