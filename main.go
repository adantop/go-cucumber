package main

import (
	"fmt"
	"github.com/adantop/cucumber/concurrency"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("main: start")

	//concurrency.BasicConcurrency()
	//concurrency.AtomicAddInt()
	concurrency.MutexConcurrency()

	defer fmt.Printf("main: took %v\n", time.Since(start))
}
