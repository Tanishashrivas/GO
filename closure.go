package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// The i variable is defined inside intSeq. Since the returned function references i, it creates a closure around it.
// A closure is a function that "remembers" the environment in which it was created, even after that environment has finished execution.
// Each call to intSeq creates a new environment with its own independent i.

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	// Closures or Anonymous functions are functions without a name. They can be assigned to variables or used inline.
	f3 := func(a int, b int) int {
		return a * b
	}

	f4 := func() string {
		return "I am a standalone anonymous func"
	}

	fmt.Println(f3(5, 7), f4())
}
