package main

import (
	"math/rand"
	"time"

	j "github.com/eliferdentr/pulse/internal/jobs"
)

func main() {
	// r := gin.Default()
	store := j.NewStore()
	manager := j.NewManager(store, 13)
	jr := j.JobRequest{
		Steps:     rand.Intn(10) + 1,
		SleepMs:   200,
		TimeoutMs: 200,
	}
	manager.StartWorkers(3)
	manager.SubmitJob(jr)
	time.Sleep(3 * time.Second)

}
