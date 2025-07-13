package main

import "fmt"


func main() {
	name:="Anvesha"

	namePointer:=&name;

	fmt.Println(name)    // BEFORE CHANGE       Anvesha
	fmt.Println(&namePointer) // 0xc000058040
	fmt.Println(*namePointer) // Anvesha
	fmt.Println(namePointer)  // 0xc000012090
	fmt.Println(&name)        // 0xc000012090

	printPointer(namePointer)

	fmt.Println(name)	// AFTER CHANGE         Alice

}

func printPointer(namePointer *string){

	fmt.Println("PRINTING INSIDE FUNCTION");
	fmt.Println(&namePointer)   // 0xc000058050
	fmt.Println(*namePointer)  // Anvesha
	fmt.Println(namePointer)  // 0xc000012090

	name:=*namePointer;
	fmt.Println(&name);        // 0xc0000aa070


	// CHANGING
	*namePointer = "Alice"
	name = "Riya";
	

	fmt.Println(name); // CHANGED VALUE INSIDE FUNCTION         Riya
}
