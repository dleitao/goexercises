package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

    cont :=  make(map[string]int)
	
	for  _, word := range strings.Fields(s) {
		elem, _ := cont[word]
		cont[word] = elem + 1
	}
	
	return cont
}

func main() {
		wc.Test(WordCount)
	}

