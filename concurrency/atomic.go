package concurrency

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

// AtomicAddInt helps to understand atomic functions, in this example
// we are reading an integer from few files with pattern atomic-*.txt
// and printing the sum at the end
func AtomicAddInt() {
	fmt.Println("  AtomicAddInt: start")

	var wg sync.WaitGroup
	var sum int64

	root := os.Getenv("HOME") + "/Workspace/cucumber/testdata/"
	files := []string{"atomic-0.txt", "atomic-1.txt", "atomic-2.txt"}

	for _, f := range files {
		wg.Add(1)

		filename := root + f

		go func() {
			defer wg.Done()

			i, err := readIntFromFile(filename)
			if err != nil {
				fmt.Printf("   File: %v error: %v", filename, err)
				return
			}

			atomic.AddInt64(&sum, i)
		}()
	}

	wg.Wait()
	fmt.Printf("  AtomicAddInt: final sum %v\n", sum)
}

// readIntFromFile given filename, opens a file and reads an int
// from it, returns error if it cannot parse an int
func readIntFromFile(filename string) (int64, error) {
	var d int64
	var content []byte

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return d, err
	}

	i, err := strconv.Atoi(string(content))
	if err != nil {
		return d, err
	}

	return int64(i), nil
}
