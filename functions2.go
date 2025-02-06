// For methods, deferred functions, panic and recover and func types

package main

import "fmt"

//You can specify the type of a func for readability
type operation func(int, int) int

// eg:
// func apply(a, b int, op operation) int {}

//struct: user-defined types that allow grouping related data together.
type Rectangle struct {
	length, width float64
}

//Methods: a function that is associated with a specific type.  It uses a receiver to specify which type it belongs to.
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// func declares a function.
// (r Rectangle) before the function name is called a receiver.
// It means Area() is a method of the Rectangle struct.
// r is like 'this' in other languages, representing the current instance.

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

// If we used value receiver (r Rectangle), a copy of the struct would be modified, not the original.
// Using pointer receiver (*Rectangle) ensures that the original Rectangle object gets modified.

//Deferred Functions: executed when the surrounding function exits, regardless of where it exits. Use Case: Cleanup resources, close files, etc.
// Defer statements work in LIFO manner

func main() {
	r := Rectangle{length: 10, width: 6}
	fmt.Println("Area ", r.Area())

	r.Scale(3)
	fmt.Println("Scaled Area ", r.Area())

	//Deferred func eg:
	fmt.Println("Start")
	defer fmt.Println("Deferred func ran")
	fmt.Println("End")
}
