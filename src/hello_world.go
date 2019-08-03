// Package 包的作用的简要说明

// 包的作用的详细说明
// 包的作用的详细说明
// 包的作用的详细说明
package main


import (
	"fmt"
	"github.com/phachon/go-logger"
	"os"
)

// 定义一些常量
const Pi = 3.14159
// 枚举
const (
	Unknown = 0
	Female = 1
	Male = 2
)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
type Color int

const (
	RED Color = iota // 0
	ORANGE // 1
	YELLOW // 2
	GREEN // ..
	BLUE
	INDIGO
	VIOLET // 6
)

type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type BitFlag int
const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send // 1 << 1 == 2
	Receive // 1 << 2 == 4
)


/*
全局的变量 可以用var定义，如果在函数内部，最好使用:=定义
 */
var (
	Home = os.Getenv("HOME")
	User = os.Getenv("USER")
	Goroot = os.Getenv("GOROOT")
)

/*
main mian函数
 */
func main() {

	logger := go_logger.NewLogger()

	logger.Info("this is a info log!")
	logger.Errorf("this is a error %s log!", "format")

	sum := Sum(1, 2)
	hostname, error := os.Hostname()
	fmt.Println(hostname, error)
	if error == nil {
		fmt.Println(hostname)
	}

	fmt.Printf("hello world! %v %s\n" , sum, "aaa")
	fmt.Print("Hello:", 23)

	arr := [10]int{1,2,3}
	for i,v := range arr {
		fmt.Println(i, v)
	}



}

/*
sum 这个函数干了啥
*/
func Sum(a, b int) int {
	return a + b
}

