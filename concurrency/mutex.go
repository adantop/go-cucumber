package concurrency

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

var filename = os.Getenv("HOME") + "/Workspace/cucumber/testdata/prime-numbers.txt"

// MutexConcurrency helps to understand mutex, mutexes ensure that a
// piece of code is executed by one coroutine at a time
func MutexConcurrency() {
	fmt.Println("  MutexConcurrency: start")
	wg.Add(2)

	// fill in data
	max := 600
	work := make([]int, max)
	for i := 0; i < max; i++ {
		work[i] = i + 3
	}

	// prepare file
	cleanFile()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("  MutexConcurrency: error opening file %v\n", err)
	}

	// dividing work
	go processWork(work[max/2:], "routine-1", f)
	go processWork(work[:max/2], "routine-2", f)

	wg.Wait()

	err = f.Close()
	if err != nil {
		fmt.Printf("  MutexConcurrency: error closing file %v\n", err)
	}

	fmt.Println("  MutexConcurrency: done")
}

// processWork here we look for a prime number in the range specified
// and when found, write it to the file. To avoid having issues, we envelop
// the writing in a mutex Lock
func processWork(work []int, name string, f *os.File) {
	defer wg.Done()

	for _, v := range work {
		if isPrime(v) {
			mutex.Lock()
			{
				fmt.Printf("    processWork[%v] writting to file\n", name)
				_, err := f.Write([]byte(strconv.Itoa(v) + "\n"))
				if err != nil {
					fmt.Printf("    processWork[%v] error %v", name, err)
				}
			}
			mutex.Unlock()
		}
	}
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}

func cleanFile() {
	_ = ioutil.WriteFile(filename, []byte{}, 0644)
}