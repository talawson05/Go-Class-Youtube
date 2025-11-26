package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"strings"
	// "io"
	"os"
)

func main() {

	// Execute with: go run . a.txt b.txt c.txt
	// or: go run . *.txt
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue // move onto next file
		}

		// // CAT the file
		// if _, err := io.Copy(os.Stdout, file); err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// 	continue
		// }

		// // calculate size of file
		// data, err := ioutil.ReadAll(file)
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// 	continue
		// }
		// fmt.Println("The file has", len(data), "bytes")

		// Recreate WordCount for a file
		var lc, wc, cc int
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := scanner.Text()
			wc += len(strings.Fields(s)) // splits based on spaces/tabs
			cc += len(s)
			lc++
		}
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)

		file.Close()
	}
}