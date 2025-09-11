package main

import (
	"fmt"
	"sync"
	"time"
)

func Timer(wg1 *sync.WaitGroup) {
	defer wg1.Add(-1)
	for i := 1; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%v\n", i)
	}
}

func main() {
	var wg sync.WaitGroup
	fmt.Printf("Начало таймера:\n")

	wg.Add(1)
	go Timer(&wg)

	wg.Wait()

	fmt.Printf("Конец таймера\n")

}
