package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	/*
	s := "élite"

	fmt.Printf("%8T %[1]v %d\n", s, len(s)) // GOTCHA: shows 6 due to UTF-8 encoding
	fmt.Printf("%8T %[1]v\n", []rune(s)) // characters
	fmt.Printf("%8T %[1]v\n", []byte(s)) // UTF-8 of chacters

	hw := "Hello, World"
	hello := hw[:5]
	world := hw[7:]
	fmt.Println(hello)
	fmt.Println(world)

	phrase := "The quick brown fox"
	a := len(phrase) // 19
	b := phrase[:3] // The
	c := phrase[4:9] // quick
	d := phrase[:4] + "slow" + phrase[9:] // replaces "quick"
	phrase += "es" // now plural (copied)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(phrase)

	// Strings are passed by reference, thus aren't copied


	astring := "a string"

	// strings.Contains(astring, "g") // true
	// strings.Contains(astring, "x") // false
	// strings.HasPrefix(astring, "a") // true
	// strings.Index(astring, "string") // 2

	astring = strings.ToUpper(astring) // "A STRING"
	fmt.Println(astring)
	*/


	// Execute with: go run . foo tony < text.txt
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough arguments")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), old)
		t := strings.Join(s, new)
		fmt.Println(t)
	}
}