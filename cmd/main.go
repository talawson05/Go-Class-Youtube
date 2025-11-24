package main

import (
	"example/hello"
	"fmt"
	"os"
)

// Execute with: go run ./cmd Tony
func main() {
	// if len(os.Args) > 1 {
	// 	// fmt.Printf("Hello, %s\n", os.Args[1])
	// 	fmt.Println(hello.Say(os.Args[1]))
	// } else {
	// 	// fmt.Println("Hello, World!")
	// 	fmt.Println(hello.Say("World"))
	// }
	fmt.Println(hello.Say(os.Args[1:]))
}