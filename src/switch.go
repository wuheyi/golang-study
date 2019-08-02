package main

import "fmt"

func main() {
	intValue := 5
	switch {
	case intValue > 5:
		fmt.Println("> ", 5)
	case intValue < 5:
		fmt.Println("< ", 5)
	default:
		fmt.Println("= ", 5)
	}



}
