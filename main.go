package main

import (
	"fmt"
	"os"
)

const (
	// Only literals or compile time funcs can be consts. Read Only, concurrent safe
	constA = 1
	constB = 2 * 4 // 8
	constC = constB << 3 // 64
	constD = "a string"
	constE = len(constD) // 8
)

var (
	classVarA = 1
	classVarB = 1.00
)

// Execute with: go run . < nums.txt
// Or: cat nums.txt | go run .
func main() {
	fmt.Printf("constA: %v\n",constA)
	fmt.Printf("constB: %v\n",constB)
	fmt.Printf("constC: %v\n",constC)
	fmt.Printf("constD: %v\n",constD)
	fmt.Printf("constE: %v\n",constE)

	fmt.Printf("classVarA: %8T, %[1]v\n", classVarA)
	fmt.Printf("bclassVarB: %8T, %[1]v\n", classVarB)

	a := 2
	b := 3.01
	fmt.Printf("a: %16T, %[1]v\n", a)
	fmt.Printf("b: %16T, %[1]v\n", b)

	a = int(b)
	fmt.Printf("a: %16T, %[1]v\n", a)

	// defined as nil/zero, never null
	var sum float64
	var counter int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		counter++
	}

	if counter == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(counter))
}