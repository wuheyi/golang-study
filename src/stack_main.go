package main

import (
	"fmt"
	stackUtil "golang-study/src/stack"
)

func main(){
	stack, err := stackUtil.NewStack(5)
	if err != nil {
		panic(err.Error())
	}
	stack.Push("a")
	if data, err := stack.Pop(); err == nil {
		fmt.Println(data)
	}
}
