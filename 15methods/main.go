package main

import "fmt"

func main() {
	fmt.Println("Structs learning")
	//no inheritence in golang; No super or parent

	shinan := User{"Shinan", "shinan@gmail.com", true, 21}
	fmt.Println(shinan)
	fmt.Printf("shinan details are : %+v \n", shinan)
	fmt.Printf("shinan name is %v and email is %v.\n", shinan.Name, shinan.Email)
	shinan.GetStatus()
	shinan.NewMail()
	fmt.Printf("shinan name is %v and email is %v.\n", shinan.Name, shinan.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "newmail@gmail.com"
	fmt.Println("New mail has been set to ", u.Email)
}
