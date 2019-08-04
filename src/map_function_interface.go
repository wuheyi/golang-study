package main

import "fmt"

type obj interface{}


func main() {

	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	intArr := []obj{1,2,3,4,5}
	strArr := []obj{"aaa", "bbb", "ccc"}
	res1 := mapFunc(mf, intArr)
	res2 := mapFunc(mf, strArr)
	fmt.Println(res1)
	fmt.Println(res2)


}

func mapFunc(mf func(obj) obj, list ...[]obj) []obj {
	result := make([]obj, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}
