package main

import (
	"fmt"

	"github.com/port-scanner/pkg/portscanner"
)

func main() {
	ps := portscanner.New("127.0.0.1", 200)
	op := ps.ScanTil(450)
	fmt.Println(op)
}
