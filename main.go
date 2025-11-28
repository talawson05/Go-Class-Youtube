package main

import (
	"fmt"
)

// func fib() func() int {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b
// 		return b
// 	}
// }

// func do(d func()) {
// 	d()
// }

func main() {
	// f := fib()
	// for x:= f(); x < 100; x = f() {
	// 	fmt.Println(x)
	// }

	// f, g := fib(), fib()
	// fmt.Println(f(), f(),f(),f())
	// fmt.Println(g(), g(),g(),g())

	// // Prints values of i as 0 through 3, but all live in same address
	// for i := 0; i < 4; i++ {
	// 	v := func() {
	// 		fmt.Printf("%d @ %p\n", i, &i)
	// 	}
	// 	do(v)
	// }


	// // GOTCHA: Splitting the creation and calling of a closure, means the value is overwritten
	// // As i is a reference to i not it's value
	// // Update: no longer accurate, Go language changed in 1.22 so that variables in for loops will instantiate on every iteration
	// s := make([]func(), 4)
	// for i := 0; i < 4; i++ {
	// 	s[i] = func() {
	// 		fmt.Printf("%d @ %p \n", i, &i)
	// 	}
	// }

	// for i := 0; i < 4; i++ {
	// 	s[i]()
	// }

	// This way creates a new variable to close over, gives unique addresses
	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		i2 := i //closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p \n", i2, &i2)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}


}