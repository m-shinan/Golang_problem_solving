package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the rating for our pizza:")
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", rating)
	fmt.Printf("type of this rating is %T", rating)

}
