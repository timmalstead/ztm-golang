package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	var counter int

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		counter++

		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				counter--
				waitGroup.Done()
			}()
			var duration = time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("waiting for", duration)
			time.Sleep(duration)
		}()
	}

	fmt.Println("waiting for goroutines")
	waitGroup.Wait()
	fmt.Println("complete")
}
