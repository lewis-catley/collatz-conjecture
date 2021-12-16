package main

import (
	"fmt"
	"time"
)

func conjecture(i int) int {
	counter := 0
	for i != 1 {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = i*3 + 1
		}
		counter++
	}
	return counter
}

func main() {
	start := time.Now()
	greatestChain := 0
	greatestChainCount := 0
	for i := 1; i <= 100000000; i++ {
		count := conjecture(i)
		if count > greatestChainCount {
			greatestChain = i
			greatestChainCount = count
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Greatest chain: %d, with count: %d\n", greatestChain, greatestChainCount)
	fmt.Printf("Executed in %s\n", elapsed)
}
