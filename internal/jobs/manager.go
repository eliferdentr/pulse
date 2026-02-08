package jobs

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

/*
manager will hold these:
jobs channel
results channel
worker count
store reference
waitgroup

it's going to be a worker pool manager
*/

type Manager struct {
	Store *Store
	Jobs  chan *JobState
	WG    sync.WaitGroup
}

func NewManager(store *Store, workerCount int) *Manager {
	return &Manager{
		Store: store,
		Jobs:  make(chan *JobState, workerCount*2),
	}
}

// starts goroutine as the count of workerCount. Every worker reads from jobs channel. is going to use WaitGroup
//creates the worker pool, in short

// manager starts n amount of workers
// every worker listens to jobs channel in infinite loop and handles the incoming job
func (m *Manager) StartWorkers(workerCount int) {

	// workerCount kadar worker başlat
	// for i := 0; i < workerCount; i++
	for i := 0; i < workerCount; i++ {
		//bir worker için daha bekleyeceğimiz için wg artır
		m.WG.Add(1)

		//concurrent olarak bir iş daha yapılacak bunun için bir goroutine başlat
		go func() {
			//worker'ın işi bitince waitgroupu azaltalım
			defer m.WG.Done()

			//worker sürekli job beklemeli. job channelını dinlesin
			for job := range m.Jobs {
				m.Store.Update(job.ID, func(j *JobState) {
					j.Status = JobStatusRunning
					fmt.Println("Job UUID: " + j.ID + " is RUNNING.")
				})
				totalSteps := job.Request.Steps
				sleepMs := job.Request.SleepMs
				progress := 0
				for step := 0; step < totalSteps; step ++ {
					time.Sleep(time.Duration(sleepMs) * time.Millisecond)
					progress = (step + 1) * 100 / totalSteps
					m.Store.Update(job.ID, func (js *JobState) {
						js.Progress = progress
					})
				}
			m.Store.Update(job.ID, func (j *JobState) {
				j.Status = JobStatusDone
				fmt.Println("Job UUID: " + j.ID + " is DONE.")
			})
			}

		}()
	}
}

func (m *Manager) SubmitJob(req JobRequest) string {
	generatedId := uuid.New().String()
	newState := &JobState{
		ID: generatedId,
		Status: JobStatusQueued,
		Progress: 0,
		Request: req,
		}
	m.Store.Set(generatedId, newState)
	m.Jobs <- newState
	return generatedId
}
