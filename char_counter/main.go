package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var w, w1 sync.WaitGroup
var total int32

func count(frag string, c *chan int32) {
	defer w1.Done()
	*c <- int32(len(frag))
}

func readFileChunks(name string, c *chan int32) {

	defer w.Done()
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	const maxSz = 10000

	b := make([]byte, maxSz)
	for {
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		w1.Add(1)
		chunk := string(b[:readTotal])
		go count(chunk, c)
	}
}

func readFileLines(name string, c *chan int32) {
	fmt.Println("LINES")
	defer w.Done()
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sca := bufio.NewScanner(file)
	if err := sca.Err(); err != nil {
		fmt.Println(err)
	}
	for sca.Scan() {
		w1.Add(1)
		go count(sca.Text(), c)
	}

}
func sumAll(c *chan int32) {
	for elem := range *c {
		atomic.AddInt32(&total, elem)
	}
}

func main() {
	start := time.Now()
	c := make(chan int32, 50)
	total = 0
	w.Add(1)
	// head -c 1G </dev/urandom >randfile
	go readFileLines("rand", &c)
	// go readFileChunks("rand", &c)
	go sumAll(&c)

	w.Wait()
	w1.Wait()
	fmt.Println(total)
	elapsed := time.Since(start)
	log.Printf("Counting took %s", elapsed)
}
