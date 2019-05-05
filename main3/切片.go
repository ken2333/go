package main

import "fmt"

func main() {

	var a []string = make([]string, 2)
	a[0] = "sun"
	fmt.Println(a[0])
	a = append(a, "3")
	fmt.Println(a)
}
