package main

import (
	"fmt"
	"math"
)

const theNumber = 4
const iterations = 6

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1 ; i < iterations ; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {	
	myZz  := Sqrt(theNumber)
	yourZz := math.Sqrt(theNumber)
	if myZz == yourZz {
		fmt.Println("Wohoo is the same")
	} else {
		fmt.Println("NOT the same")
	}
	fmt.Println("my Sqrt: ", myZz)
	fmt.Println("the math.Sqrt in the standard library: ", yourZz)
}
