package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//var strValue string
	//var intValue int
	//
	//fmt.Scanln(&strValue, &intValue)
	//
	//fmt.Println(strValue, intValue)

	//reader := bufio.NewReader(os.Stdin)
	//if input, err := reader.ReadString('\n'); err == nil {
	//	fmt.Println(input)
	//}

	file, err := os.Open("./src/stack/stack.go")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader2 := bufio.NewReader(file)
	for {
		data, readerr := reader2.ReadString('\n')
		if readerr == io.EOF {
			break
		}
		fmt.Println(data)
	}

	data2, err2 := ioutil.ReadFile("./src/stack/stack.go")
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err2)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(data2))
	ioutil.WriteFile("xxx.go", data2, 0644)
















}
