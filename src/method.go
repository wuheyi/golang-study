package main

import (
	"fmt"
	"strconv"
)

type Person1 struct {
	name string
	age int
}

func (person *Person1) dispaly(school string) string {
	return person.name + "_" + strconv.Itoa(person.age) + "_" + school
}

func main() {
	xxx := Person1{"wuheyi", 18}
	fmt.Println(xxx.dispaly("cccc"))


}
