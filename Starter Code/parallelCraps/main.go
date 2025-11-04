package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numTrials := 10000
	numProcs := runtime.NumCPU()

	start := time.Now()
	ComputeHouseEdge(numTrials)
	elapsed := time.Since(start)
	fmt.Printf("Running serially took %s",elapsed)
	fmt.Println()

	start2 := time.Now()
	ComputeHouseEdgeMultiProc(numTrials,numProcs)
	elapsed2 := time.Since(start2)
	fmt.Printf("Running serially took %s",elapsed2)
	fmt.Println()

	fmt.Println(float64(elapsed2)/float64(elapsed))
}
