// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	writeData()
}
func writeData() {
	// To start, here's how to dump a string (or just
	// bytes) into a file.

	//!!!!!!!!!!!!!REST API ENTEGRATION!!!!!!!!!!!!!!!!
	//BURAYA ELLE GİRMEMEN GEREK REST APIDEN ÇEKİP YAZ!!!!!!!!!!!
	d1 := []byte(`[
		{
		 "id": 1,
		 "firstName": "Yagmur",
		 "lastName": "Ozden",
		 "balance": 300
		},
		{
		 "id": 2,
		 "firstName": "EKatina",
		 "lastName": "Milevskaya",
		 "balance": 5000
		},
		{
		 "id": 3,
		 "firstName": "Vasek",
		 "lastName": "Zalupickiy",
		 "balance": 2000
		}
	   ] `)
	err := os.WriteFile("/Users/yagmurozden/GO/users", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/Users/yagmurozden/GO/dat2")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()
}
