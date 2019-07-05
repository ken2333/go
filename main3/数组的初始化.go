package main

import "fmt"

func main() {
	var a [4]int  //数组自动初始化为0
	fmt.Println(a)
	var b  [4] int =[4] int{2,5} //未提初始值的自动化为0
	fmt.Println(b)

	c:=[4]int {5,3:10}//指定位置初始化
	fmt.Println(c)

	d:=[...]int{1,2,3,4}  //自动赋予数组初始化的长度，长度为4
	fmt.Println(len(d))
	fmt.Println(d)

	type user struct {
		name string
		age int8
	}

	e:=[...] user{              //对于负载的结构可以省略初始化标签
		{"tom",20},
	}
	fmt.Println(e)
}