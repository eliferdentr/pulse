package jobs

import (
	"context"
)
// represents the lifecycle state of a job.
type JobStatus string

const (
	JobStatusQueued    JobStatus = "queued"
	JobStatusRunning   JobStatus = "running"
	JobStatusDone      JobStatus = "done"
	JobStatusFailed   JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
	JobStatusTimeout   JobStatus = "timeout"
)

// payload received from POST /jobs.
type JobRequest struct {
	Steps     int `json:"steps"`
	SleepMs   int `json:"sleepMs"`
	TimeoutMs int `json:"timeoutMs"`
}

// represents the outcome of a completed job.
type JobResult struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}
// JobState represents the lifecycle of a job in memory.
type JobState struct {
	ID       string            `json:"id"`
	Status   JobStatus         `json:"status"`
	Progress int               `json:"progress"`
	Result   *JobResult        `json:"result,omitempty"` //optional, only set when job is done
	Cancel   context.CancelFunc `json:"-"` // Not serialized, used to cancel the job
	Request JobRequest `json:"-"`
}