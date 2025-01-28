package main

import "fmt"

// Functions : using func keyword
func add(a int, b int) int {
	return a + b
}

// multiple return value functions, often used for errors
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero, seriously bro? 😐")
	}

	return a / b, nil
}

func two() (int, int) {
	return 1, 2
}

// Variadic functions: they can accept a variable number of arguments.
func sum(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

// HOF(Higher-Order Functions): Functions that take or return other functions.
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func applyOperation2(a, b int, op func(int, int) (int, error)) (int, error) {
	return op(a, b)
}

func main() {
	c := add(5, 7)
	fmt.Println(c)

	//   fmt.Println(two())
	//   a,b:= two()
	//   fmt.Println(b,a)

	div2, err2 := divide(6, 0)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(div2)
	}

	// Variadic
	arr1 := []int{5, 4, 7, 3}

	f1 := sum(1, 2, 3, 4)
	f2 := sum(arr1...)

	fmt.Println(f1, f2)

	//HOF
	f3 := applyOperation(45, 5, add)
	f4, err := applyOperation2(45, 5, divide)

	fmt.Println(f3, f4, err)

}
