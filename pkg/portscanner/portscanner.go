package portscanner

import (
	"fmt"
	"net"
	"time"
)

type Scanner struct {
	domain string
}

func New(domain string) Scanner {
	return Scanner{domain}
}

func (s Scanner) ScanTo(toPort int) (openPorts []int) {
	jobs := make(chan int)
	results := make(chan int)

	for i := 0; i < toPort; i++ {
		go worker(s.domain, jobs, results)
	}

	go func() {
		for i := 1; i <= toPort; i++ {
			jobs <- i
		}
	}()

	for i := 1; i <= toPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(jobs)
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
