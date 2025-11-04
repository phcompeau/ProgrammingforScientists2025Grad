package main


// ComputeHouseEdgeMultiproc takes an integer numTrials as well as an integer numProcs and returns an estimate of the house edge of craps (or whatever binary game) played over numTrials simulated games, distributed over numProcs processors.
func ComputeHouseEdgeMultiproc(numTrials, numProcs int) float64 {
	c := make(chan int, numProcs)

	count := 0 
	for i:= 0; i < numProcs; i++ {
		if i < numProcs -1 {
			go TotalWinOneProc(numTrials/numProcs,c)
		} else {
			go TotalWinOneProc(numTrials/numProcs+numTrials%numProcs,c)
		}
	}

	for i:=0; i < numProcs; i++ {
		count += <-c
	}

	return float64(count)/float64(numTrials)
}

// TotalWinOneProc
// Input: numTrials as an integer, and an integer channel
// Output: doesn't return anything, but it enters the total amount won in numTrials games of simulated craps into the channel
func TotalWinOneProc(numTrials int, c chan int) {
	count := 0 

	for i:=0; i < numTrials; i++ {
		outcome := PlayCrapsOnce()
		if outcome {
			count++
		} else {
			count--
		}
	}
	c <- count
}