/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Dispatcher Implementation
 *
 */

package watch

import (
	"github.com/sirupsen/logrus"
)

// WorkerQueue is a buffered channel of these channel type
// Each Worker/Job has an unbuffered channel to listen on
var WorkerQueue chan chan JobRequest

// InitDispatcher initializes the dispatcher
func InitDispatcher(numWorkers int, callback Callback) {

	// Initialize the WorkerQueue
	WorkerQueue = make(chan chan JobRequest, numWorkers)

	// Create the Workers and Run them
	for idx := 0; idx < numWorkers; idx++ {

		worker := CreateWorker(idx+1, WorkerQueue, callback)

		worker.Run()

		logrus.Println("Started Worker", idx+1)

	}

}

// AssignJob schedules Jobs to Workers
func AssignJob(job JobRequest) {

	// Get an Idle Worker channel
	workerChan := <-WorkerQueue

	// Assign Worker the Job to work on
	workerChan <- job

	logrus.Printf("Assigned Job: %d to Worker", job.JobID)

}

// RunDispatcher runs the dispatcher
func RunDispatcher() {
	queue := JobQueue

	logrus.Println("Run Dispatcher")

	go func() {
		// Run Forever
		for job := range queue {
			// Assign the Job to a Worker
			go AssignJob(job)
		}
	}()

}
