/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Watcher Implementation
 *
 */

package watch

import (
	log "github.com/sirupsen/logrus"
)

// JobRequest hold the Job
type JobRequest struct {
	JobID int
	Arg   interface{}
}

// JobQueue : All jobs get queued here
var JobQueue chan JobRequest

// WatcherInit intializes the Watcher
func WatcherInit(numJobs int) {

	// Initialize the Job-Q with configured number of Jobs permitted
	JobQueue = make(chan JobRequest, numJobs)
	log.Printf("Created JobQueue: %d\n", numJobs)

	return
}

// AddJob adds a Job to the worker queue
func AddJob(id int, arg interface{}) {
	// Filter Job Requests, only interested ones get Queued

	log.Printf("Before Job create, id: %d\n", id)
	// Create the JobRequest
	job := JobRequest{
		JobID: id,
		Arg:   arg,
	}

	log.Printf("Before Job queued, id:%d\n", id)
	// Push it to the JobQueue channel
	JobQueue <- job

	log.Printf("Job queued, id:%d\n", id)
	return
}
