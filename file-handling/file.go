package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	fmt.Println("Files in go!!");

	content := "Checking up the making of a demo file in Go";

	file, err := os.Create("./demo-file.txt");
	checkNilError(err);

	length, err := io.WriteString(file, content);
	checkNilError(err);
	fmt.Println("length: ", length);

	readFile("./demo-file.txt");
}

func checkNilError (err error) {
	if(err != nil) {
	panic(err)
	}
}

func readFile (filename string) {
	// ioutil.ReadFile() is deprecated
	databytes, err := os.ReadFile(filename); //returns data in bytes
	checkNilError(err);

	fmt.Println(string(databytes));
}