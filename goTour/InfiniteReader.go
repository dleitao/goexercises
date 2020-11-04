package main

import 	(
"io"
"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.


func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return int(b[0]), nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		fmt.Printf("%v" , n)
		if err == io.EOF {
			break
		}
	}
}
