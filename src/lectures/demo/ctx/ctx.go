package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	var outputChannel = make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				outputChannel <- fmt.Sprintf("%v operation completed", msg)
				return
			case <-ctx.Done():
				outputChannel <- fmt.Sprintf("%v aborted", msg)
				return
			}
		}
	}()

	return outputChannel
}

func main() {
	var initialContext = context.Background()
	var fullContext, cancelContext = context.WithCancel(initialContext)

	var webServer = sampleOperation(fullContext, "webServer", 200)
	var microService = sampleOperation(fullContext, "microService", 500)
	var database = sampleOperation(fullContext, "database", 900)

MainLoop:
	for {
		select {
		case value := <-webServer:
			fmt.Println("webServer:", value)
		case value := <-microService:
			fmt.Println("microService:", value)
			fmt.Println("cancelling context")
			cancelContext()
			break MainLoop
		case value := <-database:
			fmt.Println("database:", value)
		}
	}

	fmt.Println(<-database)
}
