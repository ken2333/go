package main

import (
	"fmt"
	"strconv"
)

type article struct {
	title string
	user  //使用了之前的user
}

type user struct {
	name string
	age  int
}

/*结构体*/
func main() {

	var u user
	u.name = "赵日天"
	u.age = 12
	//	fmt.Println(u)

	var a article
	a.user = u
	a.title = "title"
	a.age = 12
	a.name = "name"
	fmt.Println(a)
	//第二个参数表示的是进制，第三个参数表示bit的数量
	fmt.Println(strconv.ParseInt("1000111", 2, 32))
	fmt.Println(strconv.ParseInt("123", 10, 32))

	//所谓的引用类型值的是slice、map、channel三种预定义类型。
	// 引用类型必须使用make函数创建.
	m := make([]int, 0, 10)
	m = append(m, 10)
	fmt.Println(m)

}
