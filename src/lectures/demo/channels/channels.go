package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			if msg == DoExit {
				fmt.Println("exit goroutine")
				control <- ExitOk
				return
			} else {
				panic("unhandled control message!")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	const NumberOfJobs = 30
	var jobChannel = make(chan Job, NumberOfJobs)
	var resultsChannel = make(chan Result, NumberOfJobs)
	var controlChannel = make(chan ControlMsg)

	go doubler(jobChannel, resultsChannel, controlChannel)

	for i := 1; i <= NumberOfJobs; i++ {
		jobChannel <- Job{data: i}
	}

	var closeGoroutine = func() {
		controlChannel <- DoExit
		<-controlChannel
		fmt.Println("program exit")
	}

	for {
		select {
		case result := <-resultsChannel:
			fmt.Println(result)
			if result.job.data == NumberOfJobs {
				fmt.Println("finished processing before timing out")
				closeGoroutine()
				return
			}
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			closeGoroutine()
			return
		}
	}
}
