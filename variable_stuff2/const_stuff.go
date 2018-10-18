package main

import "fmt"

const Pi = 3.14

// const cannot be declared with :=
const World = "世界"

func main() {
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
