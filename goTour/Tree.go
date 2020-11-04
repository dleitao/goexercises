package main

import (
"golang.org/x/tour/tree"
"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if nil == t {
		return
	}
	Walk(t.Right,ch)
	ch <- t.Value
	Walk(t.Left,ch)
	
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool{
	ch0 := make(chan int,10) 
	ch1 := make(chan int,10)
	
	go Walk(t1, ch0)
	go Walk(t2, ch1)
	
	for i := 0; 10 > i ; i++ {
		//fmt.Printf("%v - %v \n", <-ch0,<-ch1)
		if <-ch0 != <-ch1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)) )
}
