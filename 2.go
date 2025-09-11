package main

import (
	"fmt"
	"sync"
)

func Worker(wg1 *sync.WaitGroup, jobs <-chan int, result chan<- int) {
	defer wg1.Done()
	for arg := range jobs {
		arg = arg * arg
		result <- arg
	}
}

func main() {
	jobs := make(chan int, 10)
	result := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(&wg, jobs, result)
	}

	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for i := range result {
		fmt.Printf("%v\n", i)
	}

}
