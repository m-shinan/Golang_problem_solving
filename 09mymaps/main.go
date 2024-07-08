package main

import "fmt"

func main() {
	fmt.Println("maps in golang")
	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "rb")
	fmt.Println(languages)

	//loops

	for key, value := range languages {
		fmt.Printf("for key %v , value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}

}
