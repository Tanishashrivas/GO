package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// For conversions:
// strconv.Itoa(age) // Convert int → string
// strconv.FormatFloat(pi, 'f', 2, 64) // (value, format, decimals, bitSize) // float → string
// strconv.Atoi(strNum) // Convert string → int
// strconv.ParseFloat(strNum, 64) // Convert string → float

func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name) // It just takes one word (Scanln)

	var pizza1, pizza2 string
	fmt.Println("Enter the 2 pizza flavors you crave:")
	fmt.Scan(&pizza1, &pizza2) // Seperates words by spaces (Scan)

	// // In order to take full line, we use bufio
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanln()
	// fmt.Scan() does not consume the newline (\n) character after taking input.
	// bufio.Reader.ReadString('\n') expects input until it encounters a newline (\n).
	// When fmt.Scan() takes pizza1 and pizza2, it leaves a \n in the input buffer, which gets consumed immediately when reader.ReadString('\n') is called.

	fmt.Println("Enter your feedback here:")
	feedback, _ := reader.ReadString('\n') // Reads until a new line i.e until enter is pressed

	fmt.Println(name, " likes" , pizza1, " and " ,pizza2," and they think that ", feedback)

	fmt.Println("Give a rating between 1 and 5 please:");
	rating, _ := reader.ReadString('\n')
	//  the input includes the newline (\n) and carriage return (\r) characters, to remove it we use TrimSpace

	fmt.Printf("Type of rating %T, Value: %q\n", rating, rating) //%q is for double quoted string

	// Converting to number as it is string
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if( err != nil) {
		fmt.Println(err)
	}

	fmt.Printf("Rating: %f and type of rating %T", numRating, numRating)
}
