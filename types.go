package main

import (
	"fmt"
	"math"
)

// Structs: user defined type that's useful for grouping related data together.
type person struct {
	name string
	age  int
}

// Interface: An interface in Go is a type that specifies a set of method signatures but does not implement them. Any type that implements these methods satisfies the interface automatically.
type Shape interface {
	Area() float64
}

type Rect struct {
	len float64
	wid float64
}

type Circle struct {
	rad float64
}

func (r Rect) Area() float64 {
	return r.len * r.wid
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.rad
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

// Type Switch
func identifyType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an integer:", v)
	case bool:
		fmt.Println("It's a boolean:", v)
	default:
		fmt.Println("Unknown type 😐", v)
	}
}

func main() {
	// Struct
	fmt.Println(person{"Julie", 27})             // Positional fields
	fmt.Println(person{age: 22, name: "Muskan"}) // Named fields
	// fmt.Println(person{22, "muskmelon"}) // wont work as positions are diff

	// t := person{
	// 	"Kuch bhi", 31,
	// }
	// fmt.Println(t)

	dog := struct {
		name    string
		goodBoy bool
	}{
		"Tuffy", true,
	}

	fmt.Println(dog)

	// Interface
	var r Rect = Rect{len: 7, wid: 3.5}
	c := Circle{rad: 7.5}

	printArea(r)
	printArea(c)

	// Go provides type assertions and type switches for extracting and handling different types from an interface. These features are particularly useful when dealing with empty interfaces (interface{}), where the actual type is unknown at compile time
	// Type Assertions: It allows you to extract the underlying value of an interface and convert it into a concrete type.

	var mysteryBox interface{} = "Chocolate"

	chocolate, ok := mysteryBox.(string)
	// mysteryBox → The interface variable.
	// ConcreteType, string in this case → The type you are trying to extract.
	// ok → A boolean that tells whether the assertion was successful.

	if ok {
		fmt.Println("Yay! It's a chocolate:", chocolate)
	} else {
		fmt.Println("Oops! It's not a chocolate.")
	}

	num, ok := mysteryBox.(int)

	if ok {
		fmt.Println(num)
	} else {
		fmt.Println("TA Failed, skill issue🫢")
	}

	// Type Switch: A special switch statement that checks multiple possible types of an interface value.
	identifyType("Hello")
	identifyType(42)
	identifyType(true)
	identifyType(3.14)
}
