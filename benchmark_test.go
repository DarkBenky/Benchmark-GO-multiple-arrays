// benchmark_test.go
package benchmark

import (
	"sync"
	"testing"
)

const (
	arraySize    = 10000000
	numGoroutine = 4
)

// prepareArrays creates a shared array and separate arrays
func prepareArrays() ([]int, [][]int) {
	// Create shared array
	sharedArray := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		sharedArray[i] = i
	}

	// Create separate arrays
	separateArrays := make([][]int, numGoroutine)
	for i := 0; i < numGoroutine; i++ {
		separateArrays[i] = make([]int, arraySize)
		copy(separateArrays[i], sharedArray)
	}

	return sharedArray, separateArrays
}

// processingFunction simulates some work on the data
func processingFunction(val int) int {
	return val * 2
}

// BenchmarkSharedArray tests reading from a single shared array
func BenchmarkSharedArray(b *testing.B) {
	sharedArray, _ := prepareArrays()
	
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		results := make([]int, arraySize)

		for g := 0; g < numGoroutine; g++ {
			wg.Add(1)
			start := (arraySize / numGoroutine) * g
			end := start + (arraySize / numGoroutine)
			
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					results[i] = processingFunction(sharedArray[i])
				}
			}(start, end)
		}
		
		wg.Wait()
	}
}

// BenchmarkSeparateArrays tests reading from separate arrays
func BenchmarkSeparateArrays(b *testing.B) {
	_, separateArrays := prepareArrays()
	
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		results := make([]int, arraySize)

		for g := 0; g < numGoroutine; g++ {
			wg.Add(1)
			start := (arraySize / numGoroutine) * g
			end := start + (arraySize / numGoroutine)
			
			go func(g, start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					results[i] = processingFunction(separateArrays[g][i])
				}
			}(g, start, end)
		}
		
		wg.Wait()
	}
}