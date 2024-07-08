package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println(string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
