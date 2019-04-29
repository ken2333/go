package main

import "fmt"

func main() {
	//循环语句1
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	//死循环，一般用在网络编程上比较多
	for {
		fmt.Println("123")
	}
}
