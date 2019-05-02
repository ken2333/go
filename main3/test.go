package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("/dex")
	fmt.Println(err)
	fmt.Println(f)
	const sun = 10 - 3
	var s = "123"
	fmt.Println(sun)
	fmt.Println(&s)
	strconv.ParseInt()
}
