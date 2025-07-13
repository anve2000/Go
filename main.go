package main

import "fmt"


func main() {
	// var colors map[string]string;	
	colors:=make(map[string]string);

	// colors:=map[string]string{
	// 	"red":"123",
	// 	"green":"232",
	// }

	colors["yellow"] = "879"

	colors = map[string]string{
		"Blue":"977",
	}

	delete(colors, "Blue")

	fmt.Println("Colors ", colors);
		fmt.Printf("%+v", colors);

}
