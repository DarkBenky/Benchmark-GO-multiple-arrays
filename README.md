# Array Benchmark Tests

This benchmark suite compares the performance of two approaches for parallel processing on large arrays in Go. The primary goal is to evaluate how reading from a shared array versus separate arrays affects the execution time and memory allocation when the array is processed in parallel with multiple goroutines.

## Overview

- **Shared Array Benchmark**: Measures the performance of reading from a single shared array with multiple goroutines, each accessing different segments of the array concurrently.
- **Separate Arrays Benchmark**: Measures the performance of reading from separate copies of the array for each goroutine.

The benchmarks help assess whether using separate arrays for each goroutine offers any speed or memory advantage compared to sharing a single array.

### Constants

- **arraySize**: 10,000,000 elements
- **numGoroutine**: 4 goroutines used to divide and process the array concurrently.

### Functions

- **prepareArrays**: Prepares the shared array and separate arrays. It initializes an array of `arraySize` elements and copies this array to create separate arrays for each goroutine.
- **processingFunction**: Simulates a processing task by doubling the value of each element in the array. This function is used within both benchmarks to apply some computation to the array elements.

## Benchmark Results

The benchmarks were run with the following environment:
- **Processor**: AMD Ryzen 7 1700 Eight-Core Processor
- **Architecture**: x86_64
- **OS**: Linux

Results:

| Benchmark               | Iterations | Time per Iteration (ns) | Memory per Iteration (B) | Allocations per Iteration |
|-------------------------|------------|--------------------------|---------------------------|----------------------------|
| **BenchmarkSharedArray** | 68         | 17,654,827               | 80,003,888                | 10                         |
| **BenchmarkSeparateArrays** | 68    | 17,376,049               | 80,003,702                | 10                         |

### Interpretation

- **Execution Time**: Both benchmarks show similar execution times (17.65 ms for the shared array and 17.38 ms for separate arrays), with a slight advantage for the separate arrays. This indicates that creating separate arrays for each goroutine does not introduce a significant overhead compared to using a shared array.
  
- **Memory Usage**: Both approaches also show comparable memory usage, around 80 MB per iteration. This suggests that copying arrays does not increase memory overhead in a meaningful way for this benchmark size and processing function.

### Conclusion

The choice between using a shared array and separate arrays seems to have minimal impact on both performance and memory usage in this specific case. This may vary depending on the type of processing function or array size. For high-performance applications, testing specific use cases is essential to make an informed decision.

## Running the Benchmarks

To run the benchmarks, execute the following command:

```bash
go test -bench=. -benchmem
```
