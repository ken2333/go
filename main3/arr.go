package main

import "fmt"

func main() {
	var a1 = make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		a1 = append(a1, i)
	}
	fmt.Println(a1)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	x, ok := m["c"] //其中x赋予了value,ok则是一个Boolean，如果存在value，ok 是true,否则是false
	fmt.Println(x, ok)
	fmt.Println(m["a"])
	delete(m, "a")
	fmt.Println(m["a"])
}
