package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("/dex")
	fmt.Println(err)
	fmt.Println(f)
	const sun = 10 - 3
	var s = "123"
	fmt.Println(sun)
	fmt.Println(&s)

	var result = add1(3)
	fmt.Println(result)
}

func add1(num uint64) uint64 {
	if num > 1 {
		return num * add1(num-1)
	} else {
		return 1
	}
}
