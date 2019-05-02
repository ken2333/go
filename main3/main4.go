package main

import "fmt"

func main() {
	//循环语句1
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	//死循环，一般用在网络编程上比较多
	/*for {
		fmt.Println("123")
	}*/

	//循环遍历数组
	var arr []int = []int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(&arr)

	for i, n := range arr { //i表示index,n表示value
		fmt.Println(i, ":", n)
	}

}
