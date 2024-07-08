package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// var fruitList = []string{"apple", "banana", "cherry"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "grape", "kiwi")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867
	// highScores[4] = 777

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)

	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "nodejs", "python", "java", "swift"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
