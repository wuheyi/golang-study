package main

import (
	"fmt"
	"strings"
	"reflect"
)

type Any interface {}

type UnkownedType struct {
	s1, s2, s3 string
}

func (n UnkownedType) String() string {
	return n.s1 + "_" + n.s2 + "_" + n.s3
}

func (n UnkownedType) IsContained(s string) bool {
	if strings.Contains(n.s1, s) {
		return true
	}
	return false
}




func main() {
	var a float32 = 1.2
	// 通过反射修改内容
	reflect.ValueOf(&a).Elem().SetFloat(2.3)
	fmt.Println(a)

	var secret Any = UnkownedType{"aa", "bb", "cc"}

	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	fmt.Println(value)

	fmt.Println(typ)

	kind := value.Kind()
	fmt.Println(kind)

	for i := 0; i < value.NumField(); i ++ {
		fmt.Println(value.Field(i))
	}
	s1 := value.FieldByName("s1")
	fmt.Println(s1)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("a")
	ret := value.MethodByName("IsContained").Call(params)
	fmt.Println(ret[0].Bool())
	//params[0] = reflect.ValueOf("b")
	//ret = value.MethodByName("isContained").Call(params)
	//fmt.Println(ret)





}
