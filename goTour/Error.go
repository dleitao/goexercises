package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
 	Description string
}

func Sqrt(x float64) (float64, error) {

	if  0 > x {
		 return  -1, &ErrNegativeSqrt{"Negative number"}
	}

	var z float64 = 1
	for i:=0;i < 10 ;i++ {
	 	z -=  (z*z - x) / (2*z)
	}
	
	return float64(z), nil
}

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("%s", e.Description)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
