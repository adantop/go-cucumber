package main

import (
	"fmt"
	"time"

	"github.com/adantop/cucumber/concurrency"
)

func main() {
	start := time.Now()
	fmt.Println("main: start")

	//concurrency.BasicConcurrency()
	//concurrency.AtomicAddInt()
	//concurrency.MutexConcurrency()
	concurrency.Channels()

	defer fmt.Printf("main: took %v\n", time.Since(start))
}
