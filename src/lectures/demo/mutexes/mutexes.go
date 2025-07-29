package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Hits struct {
	count int
	mutex sync.Mutex
}

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func serve(waitGroup *sync.WaitGroup, hitCounter *Hits, iteration int) {
	wait()
	hitCounter.mutex.Lock()
	defer hitCounter.mutex.Unlock()
	defer waitGroup.Done()
	hitCounter.count++
	fmt.Println("server iteration", iteration)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var waitGroup sync.WaitGroup
	var hitCounter = Hits{}

	var amountOfHitsToSimulate = 20
	waitGroup.Add(amountOfHitsToSimulate)

	for i := 0; i < amountOfHitsToSimulate; i++ {
		var currentIteration = i
		go serve(&waitGroup, &hitCounter, currentIteration)
	}
	fmt.Printf("waiting for goroutines....\n\n")
	waitGroup.Wait()

	hitCounter.mutex.Lock()
	var totalHits = hitCounter.count
	hitCounter.mutex.Unlock()
	fmt.Println("total hits", totalHits)
}
