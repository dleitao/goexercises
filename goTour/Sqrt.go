#Exercise: Loops and Functions

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i:=0;i < 10 ;i++ {
	 	z -=  (z*z - x) / (2*z)
	}
	
	return float64(z)
}

func main() {
	fmt.Println(Sqrt(89562))
	fmt.Println(math.Sqrt(89562))
}

