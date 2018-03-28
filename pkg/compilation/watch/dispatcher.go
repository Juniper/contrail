/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Dispatcher Implementation
 *
 */

package watch

import (
	log "github.com/sirupsen/logrus"
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

		log.Println("Started Worker", idx+1)

	}

}

// AssignJob schedules Jobs to Workers
func AssignJob(job JobRequest) {

	// Get an Idle Worker channel
	workerChan := <-WorkerQueue

	// Assign Worker the Job to work on
	workerChan <- job

	log.Printf("Assigned Job: %d to Worker\n", job.JobID)

}

// RunDispatcher runs the dispatcher
func RunDispatcher() {

	log.Println("Run Dispatcher")

	go func() {
		// Run Forever
		for {
			select {
			case job := <-JobQueue:
				// Assign the Job to a Worker
				go AssignJob(job)
			}
		}
	}()

}
