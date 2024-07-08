package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeterTwo()
	greeter()

	result := adder(3, 5)
	fmt.Println("The result is ", result)

	proResult, myMessage := proAdder(2, 5, 6, 7, 9)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro message is: ", myMessage)
}

func greeter() {
	fmt.Println("Hello, I am a function!")
}

func greeterTwo() {
	fmt.Println("I also am a function!")
}

func adder(a, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi everyone"
}
