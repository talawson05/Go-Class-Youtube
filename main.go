package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)


func main() {
	// var a [3]int
	// var b [3]int{0, 0, 0}
	// var c [...]{0, 0, 0}
	// var d [3]int
	// d = b // Arrays are passed by value, elements are copied

	// var a []int // nil
	// var b = []int{1, 2} // initialised with values
	// a = append(a, 1) //append to nil OK
	// b = append(b, 3) // []int{1, 2, 3}
	// a = b // overwrites a
	// d := make([]int, 5) // []int{0, 0, 0, 0, 0}
	// e := a // same storage (alias)
	// e[0] == b[0] // true
	// // slices are passed by reference, no copying, updating OK

	// map[keyType]valueType
	// var m map[string]int // nil
	// p := make(map[string]int) // non-nil but empty
	// a := p["the"] // returns 0
	// b := m["the"] // same
	// // m["and"] = 1 // panic: assignment to entry in nil map
	// m = p
	// m["and"]++ // OK, same as p now
	// c := p["and"] // returns 1

	// // Maps are passed by reference, no copying, updating OK
	// // The key type must have == and != defined (not slices, maps, or funcs)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(m)
	// fmt.Println(p)
	// fmt.Println(c)

	// var m2 = map[string]int{
	// 	"and": 1,
	// 	"the": 1,
	// 	"or": 2,
	// }

	// var n2 map[string]int
	// // b2 := m2 == n2 // syntax error
	// c2 := n2 == nil // true
	// d2 := len(m2) //3
	// // e2 := cap(m2) // type mismatch

	// fmt.Println(m2)
	// fmt.Println(n2)
	// fmt.Println(c2)
	// fmt.Println(d2)


	// p3 := map[string]int{} //non-nil but empty
	// fmt.Println(p3)
	// a3 := p3["the"] // returns 0
	// fmt.Println(a3)
	// b3, ok := p3["and"] // returns 0, false (missing from map)
	// fmt.Println(b3)
	// fmt.Println(ok)

	// p3["the"]++
	// fmt.Println(p3)
	// c3, ok := p3["the"] // 1, true (exists in map)
	// fmt.Println(c3)
	// fmt.Println(ok)
	// if w3, ok := p3["the"]; ok {
	// 	// we know w3 is not the default value
	// 	fmt.Println(w3)
	// 	fmt.Println(ok)
	// }
	

	scanner := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
	
}