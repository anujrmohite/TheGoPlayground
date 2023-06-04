package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r http.Request)  {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		
	}
}

func main()  {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.Handle("/form", formHandler)
	http.Handle("/hello", helloHandler)

	fmt.Println("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}

}
