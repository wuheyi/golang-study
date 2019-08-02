package main

import "fmt"

func main() {

	mapList := map[string]int{"x" : 1, "y" : 2}
	fmt.Println(mapList)

	for k,v := range mapList {
		fmt.Println(k, v)
	}

	mapMake := make(map[string]int, 10)
	mapMake["hello"] = 1
	fmt.Println(mapMake)
	if v, ok := mapMake["hello"]; ok {
		fmt.Println(v)
	}
	delete(mapMake, "hello")
	if v, ok := mapMake["hello"]; ok {
		fmt.Println(v)
	}
	delete(mapMake, "hello")


	type MyMap map[string]int
	items := make([]MyMap, 5)
	for i, _ := range items {
		items[i] = make(MyMap, 1)
		items[i]["hello"] = 2
	}


	map1 := map[string]int{"hello": 1, "world": 2}
	map2 := map[string]int{"name": 1, "age": 2}

	mapStringIntMerge(items[0], map1)
	mapStringIntMerge(items[1], map2)

	for _, item := range items {
		fmt.Println(item)
	}

}

func mapStringIntMerge(m1 map[string]int, m2 map[string]int) {
	for k, v := range m2 {
		if _, ok := m1[k]; ok {
			m1[k] += v
		} else {
			m1[k] = v
		}
	}
}