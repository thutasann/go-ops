package howtousecomposeinterfaces

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//
// 🔹 Step 1: Define small, reusable interfaces
//

// A Job represents something to be executed.
type Job interface {
	ID() string
	Execute() error
}

// JobResult captures the result of a job
type JobResult interface {
	JobID() string
	Success() bool
	Error() error
}

// Queue handles pushing and pulling jobs
type Queue interface {
	Enqueue(Job)
	Dequeue() Job
}

// Logger logs events
type JobLogger interface {
	Log(message string)
}

// Retry allows retrying failed jobs
type Retry interface {
	Retry(job Job, attempt int) error
}

//
// 🔹 Step 2: Compose larger interfaces from small ones
//

// Processor is a composed interface:
// It can consume jobs, execute them, log, and retry.
type Processor interface {
	Queue
	JobLogger
	Retry
	Run(workerCount int)
	Results() <-chan JobResult
}

//
// 🔹 Step 3: Implement concrete types
//

// Simple job struct
type EmailJob struct {
	jobID string
	to    string
}

func (e *EmailJob) ID() string { return e.jobID }
func (e *EmailJob) Execute() error {
	// 30% chance to fail (simulate network issues)
	if rand.Intn(10) < 3 {
		return fmt.Errorf("failed to send email to %s", e.to)
	}
	fmt.Printf("📧 Sent email to %s\n", e.to)
	return nil
}

// JobResult implementation
type Result struct {
	jobID   string
	success bool
	err     error
}

func (r *Result) JobID() string { return r.jobID }
func (r *Result) Success() bool { return r.success }
func (r *Result) Error() error  { return r.err }

// Concrete Queue
type InMemoryQueue struct {
	jobs chan Job
}

func NewInMemoryQueue(size int) *InMemoryQueue {
	return &InMemoryQueue{jobs: make(chan Job, size)}
}
func (q *InMemoryQueue) Enqueue(job Job) {
	q.jobs <- job
}
func (q *InMemoryQueue) Dequeue() Job {
	return <-q.jobs
}

// Console logger
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	log.Println("[LOG]", message)
}

// Retry policy
type ExponentialBackoffRetry struct{}

func (r *ExponentialBackoffRetry) Retry(job Job, attempt int) error {
	delay := time.Duration(attempt*attempt) * time.Second
	time.Sleep(delay)
	return job.Execute()
}

// 🔹 Step 4: Full Processor Implementation
type WorkerPool struct {
	*InMemoryQueue
	*ConsoleLogger
	*ExponentialBackoffRetry
	results chan JobResult
	wg      sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		InMemoryQueue:           NewInMemoryQueue(size),
		ConsoleLogger:           &ConsoleLogger{},
		ExponentialBackoffRetry: &ExponentialBackoffRetry{},
		results:                 make(chan JobResult, size),
	}
}

// Results channel
func (wp *WorkerPool) Results() <-chan JobResult {
	return wp.results
}

// Run starts the worker pool
func (wp *WorkerPool) Run(workerCount int) {
	for i := 0; i < workerCount; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for job := range wp.jobs {
				wp.Log(fmt.Sprintf("Worker %d picked job %s", workerID, job.ID()))

				var err error
				for attempt := 1; attempt <= 3; attempt++ {
					err = job.Execute()
					if err == nil {
						wp.results <- &Result{jobID: job.ID(), success: true}
						break
					}
					wp.Log(fmt.Sprintf("Worker %d retrying job %s (attempt %d)", workerID, job.ID(), attempt))
					err = wp.Retry(job, attempt)
					if err == nil {
						wp.results <- &Result{jobID: job.ID(), success: true}
						break
					}
				}

				if err != nil {
					wp.results <- &Result{jobID: job.ID(), success: false, err: err}
				}
			}
		}(i)
	}
}

// Close waits for workers
func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

// 🔹 Step 5: Usage
func Distributed_Job_Processor_Sample() {
	rand.Seed(time.Now().UnixNano())

	processor := NewWorkerPool(10)

	// Start workers
	go processor.Run(3)

	// Enqueue jobs
	for i := 1; i <= 8; i++ {
		job := &EmailJob{
			jobID: fmt.Sprintf("job-%d", i),
			to:    fmt.Sprintf("user%d@example.com", i),
		}
		processor.Enqueue(job)
	}

	// Close when done
	go func() {
		time.Sleep(5 * time.Second)
		processor.Close()
	}()

	// Collect results
	for res := range processor.Results() {
		if res.Success() {
			fmt.Println("✅ Success:", res.JobID())
		} else {
			fmt.Println("❌ Failed:", res.JobID(), "-", res.Error())
		}
	}
}
