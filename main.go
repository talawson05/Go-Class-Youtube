package main

import (
	// "encoding/json"
	"fmt"
)

func main() {

	/*

	var s []int // is nil
	t := []int{} // is empty which is different to nil, capacity is 0 too
	u := make([]int, 5) // list of 5 zeros (0 being default of int). GOTCHA: append will add to 6th index after all the 0s
	v := make([]int, 0, 5) // empty (not nil) list but has capacity of 5
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil) // 0, 0, []int,  true, []int(nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil) // 0, 0, []int, false, []int{}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil) // 5, 5, []int, false, []int{0, 0, 0, 0, 0}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil) // 0, 5, []int, false, []int{}
	// Managing capacity is useful for memory efficiency 
	// If you know you need a large amount of memory you can create/reserve it upfront rather than try to grab it during runtime


	// JSON use case
	var a []int
	b := []int{}
	j1, _ := json.Marshal(a)
	j2, _ := json.Marshal(b)
	fmt.Println(string(j1)) // null
	fmt.Println(string(j2)) // []

	// When checking if slice is empty, don't use slice == nil
	// Use len(slice) == 0

	*/

	// Length vs Capacity

	a := [3]int{ 1, 2, 3 }
	b := a[0:1]

	fmt.Println("a = ", a) // a =  [1 2 3]
	fmt.Println("b = ", b) // b =  [1]

	// GOTCHA - capacity was implied
	c := b[0:2]
	fmt.Println("c = ", c) // c =  [1 2]

	fmt.Println(len(b)) // 1
	fmt.Println(cap(b)) // 3

	fmt.Println(len(c)) // 2
	fmt.Println(cap(c)) // 3

	// new operator to set capacity explicitly
	d := a[0:1:1]
	fmt.Println("d = ", d) // d =  [1]
	fmt.Println(len(d)) // 1
	fmt.Println(cap(d)) // 1

	// e := d[0:2] // fails

	// Look at the addresses
	fmt.Printf("a[%p] = %v\n", &a, a) // array can't use fancy print
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	/*
	a[0xc0000a8018] = [1 2 3]
	b[0xc0000a8018] = [1]
	c[0xc0000a8018] = [1 2]
	*/

	// Because same address, it can overwrite
	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a) //[1 2 5] instead of [1 2 3]
	fmt.Printf("c[%p] = %[1]v\n", c) // [1 2 5]
	// IF we decalred c like
	// c := b[0:2:2]
	// appending beyond capacity triggers reallocation
	// then the value of a would not be overwritten
	// GOTCHA 2: even with capacity declared c can overwrite a if value is assigned before reallocation
	// e.g. c[0] = 9 before doing c = append(c,5). 
	// Once append has been done, it's safe to assign index 0 without overwriting

}
