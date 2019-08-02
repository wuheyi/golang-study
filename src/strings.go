package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world"
	upperStr := strings.ToUpper(str)
	splitStr := strings.Split(str, " ")
	newStr := strings.Join(splitStr, ",")
	isContainsHello := strings.Contains(str, "hello")
	strconv.Itoa(100)
	intValue, _ := strconv.Atoi("100")
	fmt.Println(upperStr, splitStr, newStr, isContainsHello, intValue)
}
