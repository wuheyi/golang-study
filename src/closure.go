package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file")) // returns: file.bmp
	fmt.Println(addJpeg("file")) // returns: file.jpeg


	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	where()
	// some more code
	where()


	start := time.Now()
	fmt.Println(addBmp("file")) // returns: file.bmp
	fmt.Println(addJpeg("file")) // returns: file.jpeg
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("this amount of time: %s\n", delta)

}
