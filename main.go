package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// channel-in will be used for sending the input across the goroutines
	//channel- out will be used for receiving the data from the goroutines
	in, out := make(chan int, 10), make(chan int, 10)
	// fmt.Println(in, out)

	wg := new(sync.WaitGroup)
	cwg := new(sync.WaitGroup)

	cwg.Add(1)
	go func() {
		defer cwg.Done()
		for v := range out {
			fmt.Println("rec", v)
		}
	}()

	// spanning goroutines to handle the incoming requests
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(i+1, in, out, wg)
	}

	// Input to be process
	job := []int{1, 2, 3, 4, 5}
	for _, v := range job {
		in <- v
	}

	close(in)
	wg.Wait()
	close(out)
	cwg.Wait()
}

func worker(i int, input, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range input {
		fmt.Println(i, v)
		time.Sleep(1 * time.Second)
		result <- v
	}

}
