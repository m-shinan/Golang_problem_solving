package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato", "tomato", "cabbage"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

}
