package watch

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func HandleTestMessages(ctx context.Context, oper int32, key, value string) {
	return
}

// TestDispatcherJobsCreation
func TestDispatcherJobsCreation(t *testing.T) {
	numGoRoutinesBefore := runtime.NumGoroutine()
	t.Logf("#goroutines before InitDispatcher(): %d\n", numGoRoutinesBefore)
	InitDispatcher(3, HandleTestMessages)
	numGoRoutinesAfter := runtime.NumGoroutine()
	t.Logf("#goroutines after InitDispatcher(): %d\n", numGoRoutinesAfter)
	diffGoRoutines := numGoRoutinesAfter - numGoRoutinesBefore
	if diffGoRoutines != 3 {
		t.Errorf("Unexpected number of go-subroutines %d", diffGoRoutines)
	}
}

// TestAssignJobs
func TestAssignJobs(t *testing.T) {
	InitDispatcher(2, HandleTestMessages)

	job1 := JobRequest{JobID: 1}
	job2 := JobRequest{JobID: 2}

	AssignJob(job1)
	AssignJob(job2)

	time.Sleep(2 * time.Second)
}

// TestWatcher
func TestWatcher(t *testing.T) {
	ctx := context.Background()

	WatcherInit(2)

	AddJob(ctx, 1, 0, "test1", "test")
	AddJob(ctx, 2, 0, "test2", "test2")

	job1 := <-JobQueue
	t.Logf("Created Job: %d\n", job1.JobID)
	job2 := <-JobQueue
	t.Logf("Created Job: %d\n", job2.JobID)

	if job1.JobID != 1 {
		t.Errorf("Unexpected JobID %d", job1.JobID)
	}
	if job2.JobID != 2 {
		t.Errorf("Unexpected JobID %d", job1.JobID)
	}
}
