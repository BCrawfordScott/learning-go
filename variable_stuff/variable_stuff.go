package main

import "fmt"

var x, y int = 1, 2

// thing will return 0 as the default value for int variables is 0 and this function returns a
// declared but unassigned variable.
func thing() (x int) {
	return
}

// func main() {
// 	// implicit typing to primitive.
// 	// We only need to declare a type with variable declaration.  Assignment sets a type for us.
// 	// ??? Because functions return values, anonymous functions must have a declared return value.
// 	var c, python, java = true, "no", 456
// 	fmt.Println(x, y, c, python, java, thing())
// }

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
