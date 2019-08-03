package main

import (
	"fmt"
	"reflect"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Interval struct {
	start int
	end   int
}


type life struct {
	work string
	study string
}

type Person struct {
	School
	name string
	age int
}

type School struct {
	location string
}




func newLife(work, study string) *life {
	m := new(life)
	return m
}


func main() {

	me := Person{School{"hz"}, "wuheyi", 1}
	fmt.Println(me.School.location)
	fmt.Println(me.location)


	mylife := newLife("xxx", "dddd")
	fmt.Println(mylife)


	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str= "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)


	intr1 := Interval{0, 3}
	intr2 := Interval{end:5, start:1}
	intr3 := Interval{end:5}
	fmt.Println(intr1)
	fmt.Println(intr2)
	fmt.Println(intr3)

	fmt.Println(reflect.TypeOf(intr1).Field(0))

}
