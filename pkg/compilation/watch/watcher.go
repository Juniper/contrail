/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Watcher Implementation
 *
 */

package watch

import (
	"context"

	"github.com/sirupsen/logrus"
)

// JobRequest hold the Job
type JobRequest struct {
	JobID     int64
	context   context.Context
	operation int32
	key       string
	value     string
}

// JobQueue : All jobs get queued here
var JobQueue chan JobRequest

// WatcherInit intializes the Watcher
func WatcherInit(numJobs int) {
	// Initialize the Job-Q with configured number of Jobs permitted
	JobQueue = make(chan JobRequest, numJobs)
	logrus.Printf("Created JobQueue: %d", numJobs)
}

// AddJob adds a Job to the worker queue
func AddJob(ctx context.Context, index int64, oper int32, key, value string) {
	// Filter Job Requests, only interested ones get Queued

	logrus.Printf("Before Job create, index: %d", index)
	// Create the JobRequest
	job := JobRequest{
		JobID:     index,
		context:   ctx,
		operation: oper,
		key:       key,
		value:     value,
	}

	// Push it to the JobQueue channel
	JobQueue <- job

	logrus.Printf("Job queued, index: %d", index)
}
