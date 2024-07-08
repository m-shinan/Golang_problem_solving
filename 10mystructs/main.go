package main

import "fmt"

func main() {
	fmt.Println("Structs learning")
	//no inhesritence in golang; No super or parent

	shinan := User{"Shinan", "shinan@gmail.com", true, 21}
	fmt.Println(shinan)
	fmt.Printf("shinan details are : %+v \n", shinan)
	fmt.Printf("shinan name is %v and email is %v.\n", shinan.Name, shinan.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
