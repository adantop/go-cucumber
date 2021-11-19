package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cwg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Channels() {
	court := make(chan int)

	cwg.Add(2)

	go player("Adan", court)
	go player("Poonam", court)

	court <- 1

	cwg.Wait()
}

func player(name string, court chan int) {
	defer cwg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)

			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)

		ball++
		court <- ball
	}
}
