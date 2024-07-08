// package main

// import (
// 	"html/template"
// 	"os"
// )

// // Data structure to pass to the template
// type PageData struct {
// 	Title   string
// 	Header  string
// 	Content string
// }

// func main() {
// 	// Define your template
// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Data to pass to the template
// 	data := PageData{
// 		Title:   "My Template",
// 		Header:  "Welcome to My Page",
// 		Content: "This is the content of my page.",
// 	}

// 	// Execute the template with the data
// 	err = tmpl.Execute(os.Stdout, data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Header  string
	Content string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:   "My Template",
			Header:  "Welcome to My Page",
			Content: "This is the content of my page.",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}) 
	

	http.ListenAndServe(":8080", nil)
}
