package main

import (
	"TaskScheduler/job"
	"TaskScheduler/scheduler"
	"TaskScheduler/trigger"
	"log"
	"time"
)

func main() {
	trigger := trigger.NewTrigger(time.Now().Add(10*time.Second), 5, "second")
	j1 := job.NewJob(func() { log.Println("I am job 1...") })
	j2 := job.NewJob(func() { log.Println("I am job 2...") })
	scheduler := scheduler.NewScheduler(trigger, j1, j2)
	scheduler.Start()
}
