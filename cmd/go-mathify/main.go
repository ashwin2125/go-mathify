package main

import (
	"fmt"
	"math"
	"time"

	"github.com/ashwin2125/go-mathify/pkg/arithmetic"
)

// runBenchmark runs the provided function 'fn' for 'numIterations' times,
// capturing various time metrics, and then outputs those metrics.
func runBenchmark(label string, fn func() error) {
	const numIterations = 1000000
	var totalTime time.Duration
	var minTime = time.Duration(math.MaxInt64)
	var maxTime time.Duration
	var timeSquares time.Duration

	for i := 0; i < numIterations; i++ {
		start := time.Now()
		if err := fn(); err != nil {
			fmt.Printf("Error in %s: %v\n", label, err)
			return
		}
		duration := time.Since(start)

		totalTime += duration
		timeSquares += duration * duration

		if duration < minTime {
			minTime = duration
		}
		if duration > maxTime {
			maxTime = duration
		}
	}

	avgTime := totalTime / numIterations
	stdDev := math.Sqrt(float64(timeSquares)/float64(numIterations) - float64(avgTime*avgTime))

	fmt.Printf("Benchmark %s:\n", label)
	fmt.Printf("  Average Time: %v\n", avgTime)
	fmt.Printf("  Min Time: %v\n", minTime)
	fmt.Printf("  Max Time: %v\n", maxTime)
	fmt.Printf("  Total Time: %v\n", totalTime)
	fmt.Printf("  Standard Deviation: %v\n", time.Duration(stdDev))
}

func main() {
	fmt.Println("Starting benchmarks...")

	// Run benchmarks for arithmetic.Add
	runBenchmark("Arithmetic Add", func() error {
		_ = arithmetic.Add(5, 3)
		return nil
	})

	// Run benchmarks for exponential.SquareRoot
	// runBenchmark("Exponential SquareRoot", func() error {
	// 	_, err := exponential.SquareRoot(25.0)
	// 	return err
	// })

	fmt.Println("Benchmarks completed.")
}
