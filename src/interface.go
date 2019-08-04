package main

import (
	"./util"
	"fmt"
)

type Animal util.Animaler


type ReadWrite interface {
	Read() bool
	Write() bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}



func main() {
	cat := new(util.Cat)
	cat.SetName("cat")
	speak(cat)

	pig := new(util.Pig)
	pig.SetName("pig")
	speak(pig)


	animals := []Animal{cat, pig}
	for _, animal := range animals {
		fmt.Println(animal.Speak())

		// 类型属于哪个类型
		if v, ok := animal.(*util.Cat); ok {
			fmt.Println(v.Speak())
		}

		// 判断是否实现了某个接口
		if v, ok := animal.(Animal); ok {
			fmt.Println(v.Speak())
		}

		// TYPE-SWITCH 判断属于哪个类型
		switch t := animal.(type) {
		case *util.Cat:
			fmt.Printf("%T\n", t)
		case *util.Pig:
			fmt.Printf("%T\n", t)
		case nil:
			fmt.Printf("nil value: nothing to check?\n")
		default:
			fmt.Printf("Unexpected type %T\n", t)
		}

	}



}

func speak(animal Animal) {
	fmt.Println(animal.Speak())
}