package jobs

import "sync"

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
	Jobs chan *JobState
	WG sync.WaitGroup
}

func NewManager (store *Store, workerCount int) *Manager {
	return &Manager{
		Store: store,
		Jobs: make(chan *JobState, workerCount * 2),
	}
}
// starts goroutine as the count of workerCount. Every worker reads from jobs channel. is going to use WaitGroup
//creates the worker pool, in short

//manager starts n amount of workers
//every worker listens to jobs channel in infinite loop and handles the incoming job
func (m *Manager) StartWorkers(workerCount int) {
	for i:= 0; i< workerCount; i++ {
		m.WG.Add(1)
		
	}
}