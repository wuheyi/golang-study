package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
)

func main() {
	slices := []int{1,2,3}
	fmt.Println(slices)


	var arr1 [7]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:5]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range



	var slice11 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice11); i++ {
		slice11[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice11); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice11[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice11))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice11))


	items := []int{10, 20, 30, 40, 50}
	for i, _ := range items {
		items[i] *= 2
	}
	fmt.Println(items)

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)


	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Println(s2)


	xxx := []int{3,2,4}
	sort.Ints(xxx)
	fmt.Println(xxx)





}

var digitRegexp = regexp.MustCompile("[0-9]+")
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}