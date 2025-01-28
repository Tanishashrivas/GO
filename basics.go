package main

// import ("fmt"
// "github.com/fatih/color")
import "fmt"

func main() {
	/* // printing stuff
	fmt.Println("Hello world!!")
	color.Blue("Hello world in blue")
	fmt.Println("2 + 28", 2 + 28)
	fmt.Println("Bun + Dosa", "Bun" + "Dosa")
	fmt.Println("8.0/3.0", 8.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	*/

	// Variables
	/*
		var a int = 5
		var b string = "muskan"
		var c = true // type inference
		d:= "Munch" //short-hand
		var e int // zero-values, "" for string, false for bool

		var(
			a = 5
			b = "str"
			c = true
			d = 3.14
			e int
		)
		fmt.Println("Values", a, b, c, d, e)

		var obj1 map[string]int= map[string]int{
			"Musk": 21,
			"Utk": 22,
		}
		var obj2 map[int]int= map[int]int{
			12: 21,
			22: 22,
		}
		var arr1 []int = []int{1,2,3}
		arr2 := []int{3,4,5}

		fmt.Println(obj1, obj2, arr1, arr2)

		obj3 := make(map[string]int)
		obj3["age"] = 22

		fmt.Println(obj3)
	*/

	//Constants
	// const pie float64 = 3.14159
	// fmt.Println(pie)

	/* // For loop
		for i:= 1; i<=5; i++ {
			fmt.Println(i)
		}

		// for as while loop
		j, ans := 1,0
		for j<5{
			ans += j
			j++
		}

		fmt.Println(ans)

		numbers := []int{10, 20, 30, 40, 50}

	    for index, value := range numbers { // to skip a value, we use _, like for _, value
	        fmt.Printf("Index: %d, Value: %d\n", index, value)
	    }

		if j < 0 {
			fmt.Println("j is -ve")
		} else if j%2 == 0 {
			fmt.Println("j is even")
		} else {
			fmt.Println("Kuch nhi hai bhai j")
		}

		if nums := 99; nums > 100 {
			fmt.Println("100 se bhi badkar")
		} else if nums < 100{
			fmt.Println("100 se choti", nums)
		}
	*/

	/*  //Arrays
	var a [4]int
	a[2] = 34

	b := [5]int{1,2,3,4,5}
	c := [...]int{3,4,2} // ... this mkes the compiler count the length for you
	d := []int{23, 45, 5: 7} // index: sets in-between values as 0
	length := len(c)

	fmt.Println(a,b,c,d,length)

	var e [3][3]int
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			e[i][j] = i+j
		}
	}

	f := [][]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Println(e,f)
	*/

	//smthg related to map
	a := map[string]int{
		"musk":  1,
		"muski": 2,
	}

	fmt.Println(a)
	fmt.Println(a["musk"])

	// value, exists
	_, b := a["muski"] // when we access a value using key, map returns 2 things, value(if the key exists) and boolean whether the key exists
	_, c := a["muskiiii"]

	fmt.Println(b) // returns true
	fmt.Println(c) // returns false
}
