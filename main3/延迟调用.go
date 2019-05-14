package main

import "fmt"

func defer1() {
	fmt.Println("延迟调用1")

}

func defer2() {
	fmt.Println("延迟调用2")

}

func main() {
	defer defer1()
	fmt.Println("1111")
	fmt.Println("@22")
	defer defer2()

}
