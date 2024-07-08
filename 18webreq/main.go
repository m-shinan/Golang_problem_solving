package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://info.cern.ch/hypertext/WWW/TheProject.html"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := (string(databytes))
	fmt.Println(content)
}
