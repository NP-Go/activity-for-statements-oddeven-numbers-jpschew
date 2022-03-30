package main

import "fmt"

func main() {
	//Insert code here
	var number int
	for {
		fmt.Println("Please enter a number: ")
		fmt.Scanln(&number)
		if number%2 == 0 {
			fmt.Printf("%d is even.\n", number)
		} else {
			fmt.Printf("%d is odd.\n", number)
		}
	}
}
