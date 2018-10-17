package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // := only available inside of a function (scoped variable assignment????)
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
