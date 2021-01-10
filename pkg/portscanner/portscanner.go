package portscanner

import (
	"fmt"
	"net"
	"time"
)

type Scanner struct {
	domain   string
	poolSize int
}

func New(domain string, poolSize int) Scanner {
	return Scanner{domain, poolSize}
}

func (s Scanner) ScanTil(toPort int) (openPorts []int) {
	workers := make(chan int, s.poolSize)
	results := make(chan int)

	for i := 0; i < cap(workers); i++ {
		go worker(s.domain, workers, results)
	}

	go func() {
		for i := 1; i <= toPort; i++ {
			workers <- i
		}
	}()

	for i := 1; i <= toPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(workers)
	close(results)

	return openPorts
}

func worker(domain string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", domain, j), 2*time.Second)
		if err != nil {
			results <- 0
			continue
		}

		results <- j
	}
}
