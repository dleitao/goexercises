package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
   img := make([][]uint8, dx)
   for i:=0; i < dx; i++ {
   		img[i] = make([]uint8, dy)
   }
   img[2][2] = 1
   return img
}

func main() {
	pic.Show(Pic)
}

