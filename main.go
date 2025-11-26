package main


func main() {

	// if a == b {
	// 	fmt.Println("a equals b")
	// } else {
	// 	fmt.Println("a is not equal to b")
	// }

	// if err := doSomething(); err != nil {
	// 	return err
	// }

	// // prints (0, 0) up to (9, 81)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("(%d, %d)\n", i, i*i)
	// }
	
	// // index only
	// for i := range myRange {
	// 	fmt.Println(i, myRange[i])
	// }

	// // index and value
	// for i, v := myRange {
	// 	fmt.Println(i, v)
	// }

	// // infinite loop with break
	// i, j := 0, 3
	// for {
	// 	i, j = i+ 50, j*j
	// 	fmt.Println(i, j)
	// 	if j > i {
	// 		break
	// 	}

	// 	// continue also exists to keep iterating
	// }
	
	// // nested loops
	// outer:  // label the outer loop
	// 	for k := range testImemsMap { // keys
	// 		for _, v := range returnedData { // values in list
	// 			if k == v.ID { // found it
	// 				continue outer // refer to label
	// 			}
	// 		}

	// 		t.ErrorF("key not found: %s", k)
	// 	}
	
	// switch a := f.Get(); a {
	// case 0, 1, 2: 
	// 	fmt.Println("underflow possible")

	// case 3, 4, 5, 6, 7, 8:

	// default:
	// 	fmt.Println("warning: overload")
	// }
	// // alternatives may be empty, and do not fall through. Break not required

	// // switch on "true"
	// a := f.Get()

	// switch {
	// case a <= 2:
	// 	fmt.Println("underflow possible")
	// case a <= 8:
	// 	// evaluated in order
	// default:
	// 	fmt.Println("warning: overload")
	// }

	// Package control visibility: capitalized starting character is exported, lower is package private
	// Multiple files in same package, can see package private things

	// A package can create a func init(){} which is implicitly called 

}