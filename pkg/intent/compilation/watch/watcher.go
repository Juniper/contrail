/*
 * Copyright 2018 - Praneet Bachheti
 *
 * Watcher Implementation
 *
 */

package watch

import (
  "fmt"
)

type JobRequest struct {
  JobID   int
  Arg     interface{}
}

// All jobs get queued here
var JobQueue chan JobRequest

func WatcherInit(numJobs int) {

  // Initialize the Job-Q with configured number of Jobs permitted
  JobQueue = make(chan JobRequest, numJobs)
  fmt.Printf("Created JobQueue: %d\n", numJobs)

  return
}

func AddJob(id int, arg interface{}) {
  // Filter Job Requests, only interested ones get Queued

  fmt.Printf("Before Job create, id: %d\n", id)
  // Create the JobRequest
  job := JobRequest {
    JobID: id,
    Arg:   arg,
  }

  fmt.Printf("Before Job queued, id:%d\n", id)
  // Push it to the JobQueue channel
  JobQueue <- job

  fmt.Printf("Job queued, id:%d\n", id)
  return
}
