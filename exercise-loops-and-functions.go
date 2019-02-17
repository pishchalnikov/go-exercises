package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	res := 1.0

	for i := 1; i <= 5; i++ {
		res = res - (res * res - x) / (2 * res)
	}

	return res
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
