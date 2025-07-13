package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 0, 2, 4}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println("N is even")
		} else {
			fmt.Println("N is odd")
		}
	}

}
