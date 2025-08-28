package main

import (
	"math/rand"
	"sort"
)

// RichnessMatrix takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.

func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	// loop through maps, add to a map for each 
	// for m in input maps; compute richness; add it to the map using the same key as in input map
	// need to initialize first... how?
	r := make(map[string]int)

	for name, freqmap := range allMaps {
		r[name] = Richness(freqmap)
	}

	return r
}

// SimpsonsMap takes an array of frequency maps as input. It returns a
// map mapping each sample name to its Simpson's index.

func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	// loop through maps, add to a map for each 
	// for m in input maps; compute simpsons index; add it to the map using the same key as in input map
	// need to initialize first... how?

	s := make(map[string]float64)

	for name, freqmap := range allMaps {
		s[name] = SimpsonsIndex(freqmap)
	}

	return s
}

// Downsample takes a frequency map and threshold as input. It returns a
// frequency map sampled down to m total samples.

func DownSample(freqMap map[string]int, threshold int) map[string]int {
	// check if input is actually bigger than threshold
	// grab random sample of threshold values
	// first: permute ALL values (while maintaining frequencies)
	// second: select threshold values off the top
	// return downsampled map
	total := SampleTotal(freqMap)
	if total < threshold {
		panic("DownSample() called on a frequency map with less than threshold total values")
	}

	allKeys := make([]string, 0, total)
	for key, count := range freqMap {
		for i := 0; i < count; i++ {
			allKeys = append(allKeys, key)
		}
	}

	perm := rand.Perm(total)

	newMap := make(map[string]int)

	for i := 0; i < threshold; i++ {
		key := allKeys[perm[i]]
		newMap[key]++
	}

	return newMap

}

// DownsampleMaps takes a map of frequency maps and threshold as input. It returns a map of all
// frequency map sampled down to m total samples.

func DownSampleMaps(allMaps map[string]map[string]int, threshold int) map[string]map[string]int  {
	// 
	newMaps := make(map[string]map[string]int)

	for key, freqMap := range allMaps {
		newMap := DownSample(freqMap, threshold)
		newMaps[key] = newMap
	}
	return newMaps
}


// BetaDiversityMatrix takes a map of frequency maps along with a distance metric
// ("Bray-Curtis" or "Jaccard") as input.
// It returns a slice of strings corresponding to the sorted names of the keys
// in the map, along with a matrix of distances whose (i,j)-th
// element is the distance between the i-th and j-th samples using the input metric.

func BetaDiversityMatrix(allMaps map[string](map[string]int), distMetric string) ([]string, [][]float64) {
	// loop through maps, add to a map for each 
	// for m in input maps; compute beta diversity; add it to the table of beta diversity
	// need to initialize first... how?
	numSamples := len(allMaps)
	sampleNames := make([]string, 0)
	for name := range allMaps {
		sampleNames = append(sampleNames, name)
	}
	sort.Strings(sampleNames)

	mtx := InitializeSquareMatrix(numSamples)
	
	for i := 0; i < numSamples; i++ {
		for j := 0; j < numSamples; j++ {
			if distMetric == "Bray-Curtis" {
				d := BrayCurtisDistance(allMaps[sampleNames[i]], allMaps[sampleNames[j]])
				mtx[i][j] = d
				mtx[j][i] = d
			} else if distMetric == "Jaccard" {
				d := JaccardDistance(allMaps[sampleNames[i]], allMaps[sampleNames[j]])
				mtx[i][j] = d
				mtx[j][i] = d
			} else {
				panic("Error: Invalid distance metric name given to BetaDiversityMatrix().")
			}
		}
	}
	return sampleNames, mtx
}
	







