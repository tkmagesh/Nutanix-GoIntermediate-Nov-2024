package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Spinning up %d goroutines.. Hit ENTER to start....\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1)     // increment the counter by 1
		go fn(id, wg) // scheduling the f1() execution through the scheduler to be scheduled to execute in future
	}
	wg.Wait() //block until the wg counter becomes 0 (default = 0)
	fmt.Println("Done!")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
