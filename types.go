package main

import ("fmt"
 "math")

// Structs: user defined type that's useful for grouping related data together.
type person struct {
	name string
	age int
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
	return r.len*r.wid
}

func (c Circle) Area() float64 {
	return 2*math.Pi*c.rad
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	// Struct
	fmt.Println(person{"Julie", 27})  // Positional fields
	fmt.Println(person{age: 22, name: "Muskan"}) // Named fields
	// fmt.Println(person{22, "muskmelon"}) // wont work as positions are diff

	// t := person{
	// 	"Kuch bhi", 31,
	// }
	// fmt.Println(t)
	
	dog := struct{
		name string
		goodBoy bool
	}{
		"Tuffy", true,
	}

	fmt.Println(dog)

	// Interface
	var r Rect = Rect{len: 7, wid: 3.5,}
	c := Circle{rad: 7.5}

	printArea(r)
	printArea(c)
	
}