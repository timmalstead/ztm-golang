package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandIntChannel(min, max int) <-chan int {
	var outputChannel = make(chan int, 3)

	go func() {
		for {
			outputChannel <- rand.Intn(max-min+1) + min
		}
	}()

	return outputChannel
}

func generateRandIntChannelN(count, min, max int) <-chan int {
	var outputChannel = make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			outputChannel <- rand.Intn(max-min+1) + min
		}
		close(outputChannel)
	}()

	return outputChannel
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var randIntChannel = generateRandIntChannel(1, 100)
	fmt.Println("Generating random numbers")
	fmt.Println(<-randIntChannel)
	fmt.Println(<-randIntChannel)
	fmt.Println(<-randIntChannel)
	fmt.Println(<-randIntChannel)
	fmt.Println(<-randIntChannel)

	var randIntRangeChannel = generateRandIntChannelN(2, 1, 10)

	fmt.Println("Generating pre-set amount of random numbers")
	for i := range randIntRangeChannel {
		fmt.Println(i)
	}

	var randIntRangeChannel2 = generateRandIntChannelN(4, 1, 10)

	fmt.Println("Generating pre-set amount of random numbers")
	for {
		var number, open = <-randIntRangeChannel2

		if !open {
			break
		}
		fmt.Println(number)
	}
}
