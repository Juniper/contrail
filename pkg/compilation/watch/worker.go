/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Dispatcher Implementation
 *
 */

package watch

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Worker holds worker object
type Worker struct {
	WorkerID    int
	JobChan     chan JobRequest
	ExitChan    chan bool
	WorkerQueue chan chan JobRequest
}

// CreateWorker creates a New Worker
// - Create a JobRequest Channel to listen on
// - Create an ExitChan to terminate
// - Add self to the WorkerQueue so we get JobRequests
func CreateWorker(id int, workerQueue chan chan JobRequest) Worker {

	worker := Worker{
		WorkerID:    id,
		JobChan:     make(chan JobRequest),
		WorkerQueue: workerQueue,
		ExitChan:    make(chan bool),
	}

	return worker
}

// Run runs a worker
func (g *Worker) Run() {
	go func() {
		for {
			// Be part of the Worker Queue
			g.WorkerQueue <- g.JobChan

			select {

			case job := <-g.JobChan:
				// Received a Job Request, process it
				log.Printf("Worker: %d, Received job request %d\n", g.WorkerID, job.JobID)
				time.Sleep(5 * time.Second)
				log.Printf("Worker: %d, Slept for 5 seconds\n", g.WorkerID)

			case <-g.ExitChan:
				log.Printf("Worker: %d exiting\n", g.WorkerID)
				return

			}
		}
	}()
}

// Exit quits a worker routine
func (g *Worker) Exit() {
	go func() {
		g.ExitChan <- true
	}()
}
