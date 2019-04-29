package main //main包是一个主函数，作为一个入口的标记

import "fmt" //引入包

func main() {
	fmt.Println("hello world golang")
	fmt.Println("hello world golang")

	/* 显式类型定义*/
	const var_name  string="www.baidu.com"

	/*隐式类型定义*/
	const www  = "www.gaolang.com"

	/*常量可以作为枚举使用*/
	const   (
		sun=1
		java=2
		c=3
		cpp=4
	)

	/*iota，特殊的常量，可以被编译器修改的常量,通常用于
	定义变量，下面的a=0,b就变成了1，d变成了2,以此类推
	*/
	const (
		a=iota
		b
		d
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)


}

func area()  {

	const leng int  =10
	const with  =2
	fmt.Println(leng* with)
}
