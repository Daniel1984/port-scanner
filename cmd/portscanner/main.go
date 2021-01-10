package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	workers := make(chan int, 100)
	results := make(chan int)
	openPorts := []int{}

	for i := 0; i < cap(workers); i++ {
		go worker(workers, results)
	}

	go func() {
		for i := 1; i <= 500; i++ {
			workers <- i
		}
	}()

	for i := 1; i <= 500; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(workers)
	close(results)

	fmt.Println(openPorts)
}

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		_, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", j), 2*time.Second)
		if err != nil {
			results <- 0
			continue
		}

		results <- j
	}
}
