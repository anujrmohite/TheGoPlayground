package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj45fd"

func main() {
	fmt.Println("Handling URL in GOlang")
	fmt.Println(myurl)

	//parsing URL

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	q_params := result.Query()
	fmt.Println("Types of Params is :")
	fmt.Println(q_params)

	partsOfURL := &url.URL{
		Scheme: "https",
		
	}
}
