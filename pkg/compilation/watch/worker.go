/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Dispatcher Implementation
 *
 */

package watch

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// Callback from Watch functions
type Callback func(ctx context.Context, oper int32, key, newValue string)

// Worker holds worker object
type Worker struct {
	WorkerID    int
	JobChan     chan JobRequest
	ExitChan    chan bool
	WorkerQueue chan chan JobRequest
	Callback    Callback
}

// CreateWorker creates a New Worker
// - Create a JobRequest Channel to listen on
// - Create an ExitChan to terminate
// - Add self to the WorkerQueue so we get JobRequests
func CreateWorker(id int, workerQueue chan chan JobRequest, callback Callback) Worker {

	worker := Worker{
		WorkerID:    id,
		JobChan:     make(chan JobRequest),
		WorkerQueue: workerQueue,
		ExitChan:    make(chan bool),
		Callback:    callback,
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
				logrus.Infof("Worker: %d, Received job request %d", g.WorkerID, job.JobID)

				g.Callback(job.context, job.operation, job.key, job.value)

				d := 10 * time.Millisecond
				time.Sleep(d)
				logrus.Debugf("Worker: %d, Slept for %s", g.WorkerID, d)

			case <-g.ExitChan:
				logrus.Infof("Worker: %d exiting", g.WorkerID)
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
