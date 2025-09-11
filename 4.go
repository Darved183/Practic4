package main

import (
	"fmt"
	"net/http"
	"sync"
)

var jobs = make(chan string)
var result = make(chan string)
var wg sync.WaitGroup

func Worker(jobs <-chan string, result chan<- string) {
	defer wg.Done()
	for arg := range jobs {
		Http, err := http.Get(arg)
		if err != nil {
			result <- err.Error()
			continue
		}
		result <- Http.Status
	}
}

func main() {

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(jobs, result)
	}

	List := []string{
		"https://github.com/",
		"https://github.com/",
		"https://github.com/",
		"https://github.com/",
		"https://github.com/",
		"https://www.youtube.com/",
		"https://github.com/",
	}

	go func() {
		for _, i := range List {
			jobs <- i
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
