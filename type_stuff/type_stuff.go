package main

import "fmt"

const Pi = 3.14

// Println will print the arguments one after another and return a new line. No interpolation.
// Printf will pring the given string and interpolate into it additional arguments.
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
