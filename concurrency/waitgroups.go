package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// SleepTime is the time that the go routines will take to finish
const SleepTime = 3 * time.Second

// BasicConcurrency shows working with sync.WaitGroup; WaitGroups are
// useful when there are task to complete but they do not need to keep
// track of data.
func BasicConcurrency() {
	var wg sync.WaitGroup
	fmt.Println("  BasicConcurrency: start")

	for i := 1; i <= 2; i++ {
		wg.Add(1) // Adding a counter to keep track of the waiting time

		name := fmt.Sprintf("Fn[%v]", i)

		go func() {
			defer wg.Done() // Once fn completes, reduce counter by 1

			time.Sleep(SleepTime)

			fmt.Printf("    %v: took %v\n", name, SleepTime)
		}()
		fmt.Printf("  BasicConcurrency: spawned %v\n", name)
	}

	wg.Wait() // Blocks until counter reaches 0
	fmt.Println("  BasicConcurrency: done")
}
