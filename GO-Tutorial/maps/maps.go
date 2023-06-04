package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)
	languages["Js"] = "JavaScript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ",languages);
	fmt.Println("JS shorts for: ",languages["Js"]);
	fmt.Println("Rb shorts for: ",languages["Rb"]);
	fmt.Println("Py shorts for: ",languages["Py"]);

	delete(languages,"Rb")

	fmt.Println("List of all languages after deleting: ",languages);
}