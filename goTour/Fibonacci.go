package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	acum, last := 0, 0
	
	return func () int {
	if 0 == acum{
		acum =1
		return acum
	}
	aux := acum
	acum = acum + last	
	last = aux
		return acum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

