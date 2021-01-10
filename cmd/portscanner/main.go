package main

import (
	"log"
	"sync"
)

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)
}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		log.Printf("port:%d\n", p)
		wg.Done()
	}
}
