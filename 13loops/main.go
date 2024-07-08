package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, "-", day)
	}

	for _, day := range days {
		fmt.Println(day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto hello
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("value is: ", rogueValue)
		rogueValue++
	}

	for rogueValue < 10 {
		if rogueValue == 5 {
			break
		}
		fmt.Println("value is: ", rogueValue)
		rogueValue++
	}

hello:
	fmt.Println("Hello world")
}
