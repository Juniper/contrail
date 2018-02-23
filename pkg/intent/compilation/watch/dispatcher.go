/*
 * Copyright 2018 - Praneet Bachheti
 *
 * Dispatcher Implementation
 *
 */

package watch

import (
  "fmt"
)

// Each Worker/Job has an unbuffered channel to listen on
// WorkerQueue is a buffered channel of these channel type
var WorkerQueue chan chan JobRequest

func InitDispatcher(numWorkers int) {

  // Initialize the WorkerQueue
  WorkerQueue = make(chan chan JobRequest, numWorkers)

  // Create the Workers and Run them
  for idx := 0; idx < numWorkers; idx++ {

    worker := CreateWorker(idx+1, WorkerQueue)

    worker.Run()

    fmt.Println("Started Worker", idx+1)

  }

}

func AssignJob(job JobRequest) {

  // Get an Idle Worker channel
  workerChan := <-WorkerQueue

  // Assign Worker the Job to work on
  workerChan <- job

  fmt.Printf("Assigned Job: %d to Worker\n", job.JobID)

}

func RunDispatcher() {

  fmt.Println("Run Dispacther")

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
