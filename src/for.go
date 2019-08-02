package main

import "fmt"

func main() {
	for idx := 0; idx < 10; idx++ {
		fmt.Println(idx)
	}

	str := "I AM 中国人"
	for pos, char := range str {
		fmt.Printf("position %d, char %c\n", pos, char)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	//
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}

	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
