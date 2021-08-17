package main

import (
	"math"
)

func main() {
	println(math.MaxInt)
	println(math.MinInt)

	println(uint64(math.MaxUint))
	println(uint64(^uint(0))) // same as
}
